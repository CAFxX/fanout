package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 int64) uint {
    v7 := GenS1{F0: (("" + "a") + "\n\t"), F1: (^(p1)), F2: 42}
    var v8 int = -2147483648
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint = 855
    var v2 float32 = float32((3.14159 / 1.5))
    v3 := make(map[string]string)
    var v4 int = 128
    v5 := make(map[int]bool)
    var v6 uint64 = 4294967295
    var v9 int8 = int8(v4)
    switch (v4 ^ (^(v4))) {
    case 42:
        for fi10 := 0; fi10 < 6; fi10++ {
            v2 = (v2 / v2)
            v5[v4] = true
            v11 := GenS2{F0: (!(false && false)), F1: 0}
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: "\n\t", F1: 0, F2: (v1 % (((v1 | v1)) | 1))}
    if ("foo bar" > ("" + (("\n\t" + "") + ("" + "hello")))) {
        v13 := GenS1{F0: ("a" + "foo bar"), F1: 1, F2: v1}
        _ = v13
    }
    var v14 int64 = 1
    fmt.Println(v9)
    var v15 string = (("hello" + "") + "\n\t")
    var v16 float32 = -2.25
    v17 := GenS3{F0: (("a" + ("foo bar" + "hello")) + ("a" + ("hello" + v15)))}
    _ = len(v5)
    v18 := []int{(v4 % ((v4) | 1)), (v4 / ((v4) | 1)), v4, (v4 * (-(v4)))}
    var v19 int8 = v9
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float64 = (0.0 * (1.5 / float64(v16)))
        _ = fi20
        _ = v21
    }
    v22 := v3[""]
    for fi23 := 0; fi23 < 19; fi23++ {
        if true {
            if ((0.0 + 3.14159) > 3.14159) {
                v5[(v4 - v4)] = false
                v24 := GenS1{F0: ((v22 + (v15 + ((("" + "hello") + v15) + v15))) + "hello"), F1: v14, F2: 0}
                _ = v24
            }
            if true {
                var v25 bool = false
                _ = v25
            } else {
                v19 -= 42
            }
        } else {
            var v26 int = fi23
            switch "a" {
            case "":
                if ((v6 % ((v6) | 1)) >= v6) {
                    var v27 int8 = (v9 / (((v19 >> ((42) & 63))) | 1))
                    var v28 uint = (v1 << (((v1 + v1)) & 63))
                    v29 := GenS3{F0: "a"}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    v19 += (v9 & (v19 ^ -107))
                    var v30 uint8 = uint8(v1)
                    v4 = 4294967295
                    _ = v30
                }
            case "\n\t":
                switch v15 {
                case "\n\t":
                    for fi31 := 0; fi31 < 16; fi31++ {
                        if (true || false) {
                            var v32 int64 = v14
                            _ = v32
                        }
                        _ = fi31
                    }
                case "a":
                    fmt.Println(v4, v14)
                case "":
                    v33 := GenS4{F0: v26}
                    _ = v33
                default:
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                }
            }
            v34 := GenS4{F0: int(v6)}
            _ = v26
            _ = v34
        }
        _ = fi23
    }
    var v35 bool = (42 < 18)
    v36 := v5[(v4 + int(v19))]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
