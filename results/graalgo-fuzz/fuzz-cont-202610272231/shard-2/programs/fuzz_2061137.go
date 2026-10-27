package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 int16) int {
    v6 := GenS1{F0: -207, F1: 42}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int8 = 42
        v9 := []int{(fi7 ^ 0), 2147483647}
        var v10 uint8 = 198
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return -2147483648
}
func genFn1(p0 int) bool {
    if (false && (!(1e10 <= 0.1))) {
        var v11 int8 = 42
        v12 := []int{(p0 - (p0 - p0)), (^(p0)), genFn0(42)}
        _ = v11
        _ = v12
    }
    v13 := []int{(p0 | (p0 - genFn0(42))), (p0 ^ (-(p0))), p0}
    v14 := GenS2{F0: p0}
    v15 := GenS1{F0: -142, F1: 1}
    var v16 bool = (117 <= 2147483648)
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return (v16 && true)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 bool = (true || true)
    var v4 int16 = 1000
    var v5 string = "\n\t"
    v17 := GenS3{F0: ""}
    v18 := []int{-1, -9223372036854775808}
    v4 -= 653
    _ = len(v2)
    var v19 float32 = float32((float64(1e10) / -2.25))
    for fi20 := 0; fi20 < 10; fi20++ {
        v4 -= (v4 + v4)
        v18 = append(v18, (fi20 / ((fi20) | 1)))
        if v3 {
            if len(v18) > 0 {
                v18[0] = 1
            }
            v3 = true
            if true {
                fmt.Println(v4, v4, v5)
            }
        } else {
            switch v5 {
            case "foo bar":
                var v21 int = (fi20 >> ((1) & 63))
                _ = v21
            case "\n\t":
                var v22 uint16 = 65535
                _ = v22
            }
        }
        _ = fi20
    }
    if v3 {
        if v3 {
            var v23 uint32 = 4294967295
            _ = v23
        } else {
            for fi24 := 0; fi24 < 19; fi24++ {
                fmt.Println(fi24, v19)
                for fi25 := 0; fi25 < 1; fi25++ {
                    switch v5 {
                    case "a":
                        var v26 string = ("\n\t" + ((v5 + v5) + (("" + v5) + (v5 + ("foo bar" + (v5 + v5))))))
                        _ = v26
                    }
                    _ = fi25
                }
                var v27 float32 = 0.0
                _ = fi24
                _ = v27
            }
        }
    } else {
        var v28 bool = false
        var v29 string = ""
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        v19 += 1.5
        _ = fi30
    }
    if (3.14159 >= (3.14159 / 0.1)) {
        var v31 uint = 762
        v5 = (v5 + v5)
        if ("a" >= (v5 + v5)) {
            var v32 float32 = float32((0.1 - float64(1.5)))
            _ = v32
        } else {
            v33 := GenS3{F0: ("\n\t" + ((v5 + "hello") + (v5 + ("hello" + ("a" + "hello")))))}
            _ = v33
        }
        _ = v31
    }
    if false {
        if (!v3) {
            v4 += v4
            var v34 uint16 = 65535
            _ = v34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
