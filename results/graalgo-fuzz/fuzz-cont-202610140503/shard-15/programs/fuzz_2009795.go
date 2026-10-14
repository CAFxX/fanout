package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint = 384
    v3 := make(map[string]bool)
    var v4 int16 = 1000
    var v5 uint64 = 2
    var v6 int32 = -2147483648
    v2 -= v2
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 bool = ("\n\t" == "foo bar")
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: ((("foo bar" + ("a" + ("hello" + ("foo bar" + "hello")))) + ("" + ("a" + "a"))) + ("foo bar" + ("\n\t" + "hello"))), F1: 230}
        v10 := GenS1{F0: "foo bar", F1: 0}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                v5 += v5
                if true {
                    var v13 bool = (0.1 < float64(float32((-2.25 / float64(v1)))))
                    v14 := GenS2{F0: (v4 / ((int16(v5)) | 1)), F1: 42}
                    v4 += (-(v4))
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 uint = (^(v2))
                v17 := GenS1{F0: "", F1: 2}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi11
        }
    }
    v5 = v5
    switch 9223372036854775807 {
    case 156:
        _ = len(v3)
    case -1:
        var v18 bool = false
        _ = v18
    default:
        var v19 float32 = (v1 / v1)
        _ = v19
    }
    v2 = v2
    v20 := GenS1{F0: "hello", F1: 104}
    v4 += v4
    fmt.Println(v5)
    var v21 int64 = -9223372036854775808
    if true {
        if true {
            switch ((("a" + (("" + "foo bar") + ("foo bar" + ""))) + ("hello" + "a")) + "") {
            case "foo bar":
                v22 := []int{}
                _ = v22
            }
            if true {
                for fi23 := 0; fi23 < 2; fi23++ {
                    if false {
                        var v24 float64 = 1e10
                        var v25 int16 = 979
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS1{F0: "foo bar", F1: 42}
                        _ = v26
                    }
                    var v27 int32 = 934
                    v28 := GenS1{F0: "hello", F1: 128}
                    _ = fi23
                    _ = v27
                    _ = v28
                }
                var v29 bool = ((false || ((((0.0 + 3.14159) - float64(v1)) / 1e10) > -2.25)) || (!(false || (false || false))))
                if true {
                    var v30 int8 = -67
                    v31 := GenS3{F0: (int(v4) == 127), F1: 963, F2: 1}
                    _ = v30
                    _ = v31
                }
                _ = v29
            }
        }
    }
    var v32 uint8 = 0
    v33 := []int{676, 225}
    var v34 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
