package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -32768
    var v3 int32 = 0
    var v4 uint32 = 540
    v5 := []int{1, -455, -509}
    v6 := make(map[int]bool)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch (("" + "\n\t") + "hello") {
    case "\n\t":
        v3 += -1
    }
    v7 := GenS1{F0: "a", F1: 0}
    v4 = 579
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!false) {
        if len(v5) > 0 {
            v5[0] = 514
        }
    }
    delete(v6, 365)
    delete(v6, -470)
    v8 := v6[32767]
    v9 := GenS2{F0: "", F1: (v3 & 1)}
    v8 = (true || (!(false || v8)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
