package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 int32 = 924
    v5 := v1[127]
    fmt.Println(v5, v4)
    for fi6 := 0; fi6 < 20; fi6++ {
        v5 = (^(fi6))
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = -1
            fmt.Println(v8, fi7, fi6)
            v9 := []int{42}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        switch ("hello" + (("\n\t" + "hello") + "")) {
        case "hello":
            v2[4294967295] = ("a" + "hello")
        case "a":
            var v11 uint32 = 491
            _ = v11
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: 255}
                switch ("a" + ("\n\t" + "a")) {
                case "\n\t":
                    var v14 uint = 1
                    _ = v14
                }
                switch (^(fi10)) {
                case 4294967295:
                    if true {
                        switch (-(fi12)) {
                        case 127:
                            var v15 uint16 = 32767
                            _ = v15
                        case 2147483647:
                            v5 += fi12
                            fmt.Println(fi10)
                        case -2147483648:
                            v16 := GenS1{F0: 42}
                            _ = v16
                        }
                    }
                case 318:
                    var v17 int64 = -9223372036854775808
                    var v18 string = (("foo bar" + "\n\t") + (((("foo bar" + (("" + "hello") + "\n\t")) + "\n\t") + "foo bar") + "hello"))
                    v18 = "foo bar"
                    _ = v17
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
        }
        _ = fi10
    }
    v4 += (v4 * 2147483647)
    v19 := GenS1{F0: 114}
    v20 := GenS1{F0: 213}
    v21 := GenS2{F0: 1, F1: 368, F2: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
