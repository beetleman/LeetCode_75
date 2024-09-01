(ns beetleman.leetcode-75.reverse-words-in-a-string
  (:require [clojure.string :as str]))

(defn reverse-words [s]
  (->> (re-seq #"\w+" s)
       reverse
       (str/join " ")))
