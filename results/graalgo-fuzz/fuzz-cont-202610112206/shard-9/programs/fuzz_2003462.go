package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 uint16) uint64 {
    fmt.Println(p0)
    if false {
        p0 = (p0 - p0)
    }
    fmt.Println(p0)
    v6 := GenS1{F0: (p0 << ((1) & 63))}
    var v7 string = ("hello" + ("" + ("foo bar" + "a")))
    _ = p0
    _ = v6
    _ = v7
    return 31
}
func genFn1(p0 uint32) int32 {
    var v8 bool = true
    var v9 uint32 = 641
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{-791}
        _ = fi10
        _ = v11
    }
    v12 := []int{-9223372036854775808, -2147483648}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = v12
    _ = v8
    _ = v9
    return 250
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 1
    var v3 float32 = -2.25
    var v4 int = 1
    v5 := make(map[int]int)
    v13 := GenS1{F0: 65535}
    v2 = v2
    if ((!false) && (-1 <= 42)) {
        v2 += (v2 + v2)
        fmt.Println(v2, v4)
    }
    var v14 float64 = float64(v3)
    fmt.Println(v14, v2)
    var v15 string = "hello"
    var v16 uint64 = (v2 >> ((876) & 63))
    var v17 int16 = -32768
    if (2147483648 != 368) {
        if ("\n\t" != "a") {
            var v18 uint64 = 993
            _ = len(v1)
            var v19 uint32 = 532
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 uint8 = 42
                v22 := GenS2{F0: 284, F1: true}
                var v23 float64 = 1e10
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        v24 := GenS3{F0: 255, F1: 42, F2: 255}
        _ = v24
    } else {
        _ = len(v5)
    }
    v25 := []int{(v4 ^ v4), (v4 / (((v4 - v4)) | 1))}
    var v26 string = "\n\t"
    if ("a" != "a") {
        switch ((((v15 + ((v26 + (("foo bar" + v26) + ("a" + (("\n\t" + "\n\t") + (v15 + v26))))) + v26)) + "hello") + "foo bar") + (v15 + ("hello" + v26))) {
        case "foo bar":
            for fi27 := 0; fi27 < 14; fi27++ {
                switch 127 {
                case 1:
                    v2 = (v2 << ((309) & 63))
                case 127:
                    v28 := GenS3{F0: 1, F1: 42, F2: 1}
                    _ = v28
                case 32767:
                    v14 -= (v14 - (((v14 * v14) / 0.1) / float64((((3.14159 - 0.0) * 1.5) - v3))))
                }
                var v29 bool = ("a" == v26)
                var v30 uint = 65535
                _ = fi27
                _ = v29
                _ = v30
            }
        case "\n\t":
            v2 = genFn0(65535)
        case "hello":
            if true {
                var v31 int32 = -1
                var v32 int = (v4 * (v4 * 42))
                _ = v31
                _ = v32
            } else {
                v33 := GenS4{F0: v15, F1: 42}
                for fi34 := 0; fi34 < 4; fi34++ {
                    v26 = (v26 + "foo bar")
                    _ = len(v25)
                    v35 := GenS3{F0: 0, F1: int8(fi34), F2: 484}
                    _ = fi34
                    _ = v35
                }
                _ = v33
            }
        }
        var v36 float32 = -2.25
        fmt.Println(v4)
        _ = v36
    }
    var v37 string = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
