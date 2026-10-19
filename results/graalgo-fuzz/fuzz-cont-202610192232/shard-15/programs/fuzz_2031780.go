package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int; F2 string }
func genFn0(p0 bool, p1 uint, p2 uint) int64 {
    v4 := []int{}
    var v5 int = -94
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 string = "hello"
    var v2 string = ((("" + "") + ("\n\t" + v1)) + v1)
    v3 := make(map[string]uint32)
    v1 = (v2 + v1)
    if (128 != 177) {
        var v6 int8 = 127
        v1 = "\n\t"
        _ = v6
    }
    var v7 float32 = 1.5
    v8 := GenS1{F0: 946, F1: ("a" + v1), F2: 475}
    var v9 string = "\n\t"
    var v10 int8 = 0
    v9 = v1
    fmt.Println(v10, v7, v7)
    switch (((v9 + "\n\t") + "") + "hello") {
    case "hello":
        switch v1 {
        case "hello":
            var v11 int64 = genFn0(false, 65535, 779)
            _ = v11
        }
    }
    var v12 bool = false
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            if (255 >= 2) {
                switch fi14 {
                case -2147483648:
                    v15 := GenS1{F0: 2, F1: v1, F2: 113}
                    _ = v15
                case -9223372036854775808:
                    v9 = v2
                default:
                    var v16 string = ("" + "")
                    _ = v16
                }
            }
            var v17 int64 = 1
            if (!v12) {
                var v18 uint16 = 255
                v19 := GenS2{F0: 668, F1: -1}
                _ = v18
                _ = v19
            }
            _ = fi14
            _ = v17
        }
        v20 := []int{}
        _ = fi13
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            if v12 {
                v23 := GenS1{F0: 2, F1: "hello", F2: 65535}
                _ = v23
            } else {
                if false {
                    if v12 {
                        v3["hello"] = 2147483648
                    }
                } else {
                    var v24 uint8 = 0
                    var v25 uint8 = v24
                    _ = v24
                    _ = v25
                }
            }
            v26 := GenS3{F0: -1}
            _ = fi22
            _ = v26
        }
        if false {
            v27 := GenS4{F0: ("a" + v2)}
            v28 := GenS5{F0: fi21, F1: fi21, F2: v1}
            v29 := []int{fi21, -9223372036854775808, (-(fi21))}
            _ = v27
            _ = v28
            _ = v29
        } else {
            v12 = (0.0 <= (-2.25 / ((0.1 + 0.0) * -2.25)))
            fmt.Println(fi21, fi21)
        }
        v30 := GenS5{F0: fi21, F1: (fi21 + fi21), F2: "a"}
        _ = fi21
        _ = v30
    }
    v31 := GenS4{F0: (v9 + (("hello" + v9) + v1))}
    v9 = (v1 + ("a" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v7
    _ = v8
    _ = v9
}
