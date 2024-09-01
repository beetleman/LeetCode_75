(ns beetleman.leetcode-75.merge-strings-alternately
  (:require [clojure.string :as str]))

(defn merge-alternately [word1 word2]
  (let [spaces (repeat "")
        lenght (max (count word1)
                    (count word2))]
    (str/join (interleave (take lenght (concat word1 spaces))
                          (take lenght (concat word2 spaces))))))
