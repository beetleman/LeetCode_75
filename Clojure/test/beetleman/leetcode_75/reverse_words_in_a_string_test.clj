(ns beetleman.leetcode-75.reverse-words-in-a-string-test
  (:require [beetleman.leetcode-75.reverse-words-in-a-string :as sut]
            [clojure.test :as t :refer [is deftest]]))

(deftest reverse-words-test
  (doseq [{:keys [input output]} [{:input  "the sky is blue"
                                   :output "blue is sky the"}
                                  {:input  "  hello world  "
                                   :output "world hello"}
                                  {:input  "a good   example"
                                   :output "example good a"}]]
    (is (= output (sut/reverse-words input)))))
