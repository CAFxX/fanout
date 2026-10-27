package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 uint8, p1 bool) string {
    v5 := []int{128, 358, -1, 153}
    var v6 float32 = -2.25
    if false {
        v7 := GenS1{F0: false}
        v8 := GenS2{F0: -562, F1: 82, F2: 876}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v9 uint16 = 42
    v10 := GenS3{F0: "foo bar"}
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            if (3.14159 < (-2.25 / -2.25)) {
                v9 += (v9 | v9)
                v9 = 65535
            } else {
                if true {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v9 += 65535
                        _ = fi12
                    }
                    var v13 int16 = 32767
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v13 = v13
                        if true {
                            v15 := GenS1{F0: true}
                            v16 := GenS1{F0: false}
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS4{F0: 1}
                        _ = fi14
                        _ = v17
                    }
                    _ = v13
                }
                switch (-(fi11)) {
                case 42:
                    v18 := GenS3{F0: ("a" + "foo bar")}
                    _ = v18
                }
                var v19 uint = 917
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                _ = len(v2)
                _ = fi20
            }
        } else {
            v21 := GenS5{F0: 20, F1: -102, F2: fi11}
            _ = v21
        }
        switch "\n\t" {
        case "hello":
            var v22 uint8 = 42
            _ = v22
        case "a":
            var v23 int16 = 42
            _ = v23
        case "":
            v24 := GenS3{F0: (("\n\t" + "a") + "a")}
            _ = v24
        }
        _ = fi11
    }
    var v25 uint64 = 42
    v26 := GenS3{F0: "\n\t"}
    v25 = (v25 & v25)
    v1[2147483647] = (("hello" + "\n\t") + "foo bar")
    v9 -= v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
