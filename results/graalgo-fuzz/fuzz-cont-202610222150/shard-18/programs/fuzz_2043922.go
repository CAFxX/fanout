package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
func genFn0() bool {
    if (477 <= 1) {
        var v4 float64 = (-2.25 / (0.1 / 1e10))
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 uint8 = 127
            for fi7 := 0; fi7 < 1; fi7++ {
                switch fi7 {
                case 32767:
                    v8 := []int{(fi7 / ((42) | 1)), (fi5 ^ (-(fi7))), 42, fi5}
                    _ = v8
                }
                _ = fi7
            }
            var v9 string = "hello"
            _ = fi5
            _ = v6
            _ = v9
        }
        _ = v4
    }
    if true {
        if true {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 int = 32767
                v11 -= fi10
                _ = fi10
                _ = v11
            }
            switch "foo bar" {
            case "a":
                switch -1 {
                case 1:
                    var v12 int32 = 0
                    _ = v12
                case -207:
                    var v13 int32 = 42
                    _ = v13
                default:
                    var v14 float64 = 1.5
                    _ = v14
                }
            case "hello":
                var v15 float64 = 0.0
                _ = v15
            case "":
                switch "foo bar" {
                case "foo bar":
                    switch (("foo bar" + "a") + ("hello" + "hello")) {
                    case "a":
                        var v16 float64 = 0.1
                        _ = v16
                    }
                case "a":
                    var v17 float32 = (1.5 / float32(1.5))
                    _ = v17
                }
            }
            if false {
                var v18 uint16 = 157
                v19 := []int{}
                switch -791 {
                case 1:
                    if true {
                        v20 := GenS1{F0: 33, F1: (("" + "\n\t") + "\n\t"), F2: "foo bar"}
                        _ = v20
                    }
                case 374:
                    var v21 uint16 = (v18 - (v18 ^ 1))
                    _ = v21
                case 42:
                    fmt.Println(v18)
                }
                _ = v18
                _ = v19
            }
        }
        var v22 int64 = 9223372036854775807
        var v23 uint16 = 929
        _ = v22
        _ = v23
    } else {
        var v24 float32 = 1.5
        v24 -= v24
        _ = v24
    }
    var v25 uint8 = 255
    fmt.Println(v25)
    _ = v25
    return false
}

func main() {
    var v1 uint = 464
    var v2 string = "a"
    var v3 uint16 = 1
    fmt.Println(v3)
    fmt.Println(v3, v2)
    var v26 uint64 = 111
    if true {
        v1 = (v1 & (v1 >> (((v1 << ((v1) & 63))) & 63)))
        var v27 uint64 = (v26 - (v26 % ((v26) | 1)))
        _ = v27
    }
    var v28 bool = (-1 != 1)
    v3 = (^(v3))
    if v28 {
        var v29 uint8 = 242
        var v30 int64 = -9223372036854775808
        _ = v29
        _ = v30
    }
    if v28 {
        var v31 int = 1
        v32 := GenS1{F0: 42, F1: "\n\t", F2: v2}
        if v28 {
            if (true || false) {
                v31 -= (v31 & v31)
                switch v2 {
                case "hello":
                    v33 := []int{v31, (v31 ^ (^(v31)))}
                    _ = v33
                case "\n\t":
                    v34 := []int{(^(v31))}
                    _ = v34
                }
                var v35 int = v31
                _ = v35
            }
            v36 := []int{(v31 / ((v31) | 1)), (v31 % (((v31 + v31)) | 1))}
            _ = v36
        }
        _ = v31
        _ = v32
    }
    var v37 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v37
}
