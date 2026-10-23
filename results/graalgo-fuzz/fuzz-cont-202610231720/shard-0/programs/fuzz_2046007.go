package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]string)
    var v3 int = 9223372036854775807
    var v4 int16 = 1000
    v5 := make(map[int]bool)
    v6 := GenS1{F0: -9223372036854775808, F1: 4294967295}
    v3 = (v3 % (((v3 - v3)) | 1))
    if true {
        switch ("foo bar" + ("\n\t" + "\n\t")) {
        case "a":
            var v7 float64 = v1
            _ = v7
        case "":
            v1 -= v1
        }
        v1 = v1
    }
    fmt.Println(v3, v3)
    v8 := GenS1{F0: 868, F1: 612}
    var v9 string = "a"
    if true {
        v4 -= 1000
    }
    fmt.Println(v4, v1)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int = 128
        v4 -= (v4 / ((v4) | 1))
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 bool = (!(!(-128 != 42)))
            if true {
                v9 = ("a" + "\n\t")
                v14 := []int{(^(fi10)), int(v4), (fi10 + fi12)}
                fmt.Println(v1, v13, v4)
                _ = v14
            } else {
                _ = len(v5)
                var v15 uint16 = 65535
                _ = v15
            }
            var v16 int16 = (v4 | 32767)
            _ = fi12
            _ = v13
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    v9 = "hello"
    v3 -= v3
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
