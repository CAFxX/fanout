package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
func genFn0(p0 uint32, p1 uint64) string {
    if (-1 == 1) {
        v4 := GenS1{F0: -1}
        var v5 int64 = -9223372036854775808
        v6 := GenS2{F0: p1, F1: 1000}
        _ = v4
        _ = v5
        _ = v6
    }
    if (94 == 2) {
        var v7 bool = false
        _ = v7
    }
    var v8 int8 = 42
    _ = p0
    _ = p1
    _ = v8
    return ("\n\t" + "a")
}

func main() {
    var v1 uint = 0
    var v2 int64 = 42
    var v3 int8 = 42
    var v9 int = -9223372036854775808
    v10 := GenS1{F0: 0}
    var v11 float32 = -2.25
    var v12 float32 = (v11 - float32(float64(v11)))
    var v13 int8 = (v3 & 1)
    v14 := GenS1{F0: 1}
    fmt.Println(v1, v11)
    var v15 int16 = 32767
    if true {
        if (true && true) {
            for fi16 := 0; fi16 < 7; fi16++ {
                v13 -= v3
                var v17 int32 = 42
                var v18 string = "foo bar"
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int32 = 82
            _ = v19
        }
        var v20 string = ""
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS3{F0: (v1 | 4294967295), F1: v3, F2: 2147483648}
            switch ("foo bar" + ("hello" + v20)) {
            case "foo bar":
                var v23 int16 = v15
                _ = v23
            case "\n\t":
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 string = ((v20 + v20) + "hello")
                    _ = fi24
                    _ = v25
                }
            case "a":
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := []int{-1, fi26, fi21}
                    var v28 uint = 2
                    var v29 int = fi21
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        var v30 int16 = v15
        var v31 float64 = 0.1
        var v32 int = v9
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 float64 = float64(3.14159)
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := []int{(fi34 % ((2147483647) | 1)), (fi34 * v9)}
        if (19 <= -1) {
            var v36 string = "\n\t"
            if true {
                switch "a" {
                case "a":
                    var v37 uint32 = 4294967295
                    _ = v37
                }
            }
            _ = v36
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v9
}
