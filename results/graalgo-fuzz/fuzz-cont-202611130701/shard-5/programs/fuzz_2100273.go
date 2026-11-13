package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64, p1 int64) string {
    p1 = p1
    v5 := []int{4294967295, 128, -9223372036854775808}
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 32767, F1: 336}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("a" + "foo bar")
}
func genFn1(p0 int, p1 int64, p2 int16) bool {
    v7 := []int{-9223372036854775808, p0}
    p0 = p0
    p2 = (p2 & -328)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 1
    v3 := []int{}
    var v4 uint16 = 65535
    var v8 string = (v1 + v1)
    _ = len(v3)
    if false {
        var v9 int = -815
        for fi10 := 0; fi10 < 18; fi10++ {
            if (42 > 781) {
                v3 = append(v3, (fi10 + fi10))
                v11 := GenS2{F0: 1, F1: 42, F2: 173}
                _ = v11
            }
            _ = fi10
        }
        v12 := GenS3{F0: -1, F1: -1, F2: 1}
        _ = v12
        _ = v9
    }
    var v13 int32 = 0
    if (false || (!false)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v14 int = 9223372036854775807
        var v15 int = (-(v14))
        var v16 bool = (0 < 255)
        _ = v14
        _ = v15
        _ = v16
    }
    v4 += v4
    v13 = (v13 >> ((701) & 63))
    v8 = (v1 + "foo bar")
    switch v1 {
    case "a":
        switch 128 {
        case -479:
            var v17 uint16 = (v4 + 2)
            _ = v17
        case 42:
            if (v1 < (("a" + v1) + "foo bar")) {
                var v18 uint32 = 181
                var v19 int8 = -128
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v18
                _ = v19
            } else {
                v20 := GenS4{F0: 0, F1: v13}
                _ = v20
            }
        }
    case "":
        var v21 uint32 = 0
        _ = v21
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v8 = ""
        var v22 uint16 = 32767
        switch ("hello" + v8) {
        case "a":
            switch (v8 + (v1 + v1)) {
            case "hello":
                v3 = append(v3, -2147483648)
            }
        case "hello":
            if ((((0.1 / float32(3.14159)) + -2.25) + float32(3.14159)) <= 0.0) {
                fmt.Println(v4)
                var v23 int = 0
                v24 := GenS4{F0: 39, F1: (v13 - (v13 / (((v13 + -160)) | 1)))}
                _ = v23
                _ = v24
            } else {
                var v25 uint8 = 127
                _ = v25
            }
        }
        _ = v22
    } else {
        var v26 int = 42
        fmt.Println(v13)
        _ = v26
    }
    v3 = append(v3, -149)
    if true {
        var v27 uint8 = 128
        for fi28 := 0; fi28 < 14; fi28++ {
            v27 = v27
            v2 += 42
            _ = fi28
        }
        v13 += (v13 + 493)
        _ = v27
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v29 int32 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
