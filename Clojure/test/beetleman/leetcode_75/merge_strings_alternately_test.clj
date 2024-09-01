(ns beetleman.leetcode-75.merge-strings-alternately-test
  (:require [beetleman.leetcode-75.merge-strings-alternately :as sut]
            [clojure.test :as t :refer [deftest is]]))

(deftest merge-alternately-test
  (doseq [{:keys [expected word1 word2]} [{:expected "apbqcr"
                                           :word1    "abc"
                                           :word2    "pqr"}
                                          {:expected "apbqrs"
                                           :word1    "ab"
                                           :word2    "pqrs"}
                                          {:word1    "abcd"
                                           :word2    "pq"
                                           :expected "apbqcd"}]]
    (is (= expected
           (sut/merge-alternately word1 word2)))))
