package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 string; F1 int; F2 uint16 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if (true && (-128 <= -128)) {
            if ((("a" < (("foo bar" + "a") + "foo bar")) || false) || false) {
                var v6 int16 = 32767
                _ = v6
            }
            if (2 < 255) {
                var v7 int = fi5
                _ = v7
            } else {
                var v8 uint64 = 42
                var v9 string = "hello"
                _ = v8
                _ = v9
            }
        } else {
            var v10 int32 = 1
            v11 := []int{(-(fi5)), fi5, (fi5 << ((0) & 63)), (fi5 & fi5)}
            _ = v10
            _ = v11
        }
        _ = fi5
    }
    var v12 uint64 = 18446744073709551615
    var v13 uint64 = 42
    var v14 int16 = 1000
    _ = v12
    _ = v13
    _ = v14
    return 361
}
func genFn1(p0 string, p1 string, p2 uint8) int {
    v15 := []int{128}
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 uint16 = 255
    var v4 uint64 = 276
    if (-1 > -128) {
        if true {
            var v16 int64 = 57
            _ = v16
        } else {
            if (308 > 880) {
                v3 -= (^(v3))
            } else {
                var v17 uint16 = v3
                v18 := GenS1{F0: -32768, F1: 0, F2: 0}
                _ = v17
                _ = v18
            }
        }
    } else {
        switch 9223372036854775807 {
        case 1:
            var v19 float64 = 1.5
            _ = v19
        }
    }
    var v20 float64 = (0.1 * 0.1)
    if true {
        switch "foo bar" {
        case "":
            v2["hello"] = v1
        case "\n\t":
            v21 := GenS1{F0: 0, F1: 9223372036854775807, F2: 1}
            _ = v21
        case "hello":
            if true {
                var v22 int64 = 9223372036854775807
                v23 := []int{2147483647}
                v23 = append(v23, genFn1("\n\t", "hello", 0))
                _ = v22
                _ = v23
            }
        default:
            switch (("" + "") + (((("" + "a") + ("\n\t" + "\n\t")) + ("hello" + "\n\t")) + "a")) {
            case "hello":
                fmt.Println(v3, v20, v20)
            }
        }
    } else {
        v4 += (v4 << ((65535) & 63))
        v24 := GenS1{F0: 42, F1: 0, F2: 0}
        var v25 uint = 16
        _ = v24
        _ = v25
    }
    if (v20 <= v20) {
        v26 := []int{}
        v20 -= v20
        _ = v26
    } else {
        v1 -= (v1 | v1)
    }
    v27 := GenS2{F0: "\n\t", F1: 2, F2: 0}
    var v28 bool = false
    switch "\n\t" {
    case "a":
        v1 = 584
    case "hello":
        delete(v2, "\n\t")
    default:
        v1 += (^(v1))
    }
    var v29 uint8 = 46
    var v30 int = 127
    if v28 {
        var v31 uint64 = 4294967295
        _ = v31
    } else {
        switch (v30 & v30) {
        case 9223372036854775807:
            fmt.Println(v20, v20, v4)
        }
        fmt.Println(v30, v29, v28)
        v32 := GenS3{F0: "hello", F1: (v30 / (((v30 + v30)) | 1)), F2: v3}
        _ = v32
    }
    v3 = 0
    if ((^(v1)) >= uint32(v29)) {
        var v33 uint16 = 65535
        v29 += 118
        for fi34 := 0; fi34 < 3; fi34++ {
            v1 = uint32(v29)
            v35 := GenS2{F0: "foo bar", F1: 37, F2: -1}
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    v36 := []int{127, -309, v30}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
