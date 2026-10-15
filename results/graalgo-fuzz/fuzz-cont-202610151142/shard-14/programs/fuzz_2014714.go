package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    v1 := []string{""}
    var v2 int = 127
    v3 := []int{799}
    var v4 uint16 = 2
    var v5 int = int(v4)
    fmt.Println(v4, v2)
    for fi6 := 0; fi6 < 18; fi6++ {
        if (false && true) {
            v2 += fi6
            v7 := GenS1{F0: ((("foo bar" + "") + "foo bar") + ("" + "\n\t")), F1: 255, F2: 18446744073709551615}
            _ = v7
        }
        var v8 bool = false
        _ = fi6
        _ = v8
    }
    v9 := GenS1{F0: "\n\t", F1: 0, F2: 666}
    v10 := GenS1{F0: "hello", F1: 42, F2: 4294967295}
    switch (v2 << ((42) & 63)) {
    case 1:
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint64 = 0
            _ = fi11
            _ = v12
        }
    case 42:
        var v13 float32 = 3.14159
        _ = v13
    }
    fmt.Println(v4, v4)
    if (2 <= v4) {
        v2 = (v2 % ((int(v4)) | 1))
    }
    var v14 int = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
