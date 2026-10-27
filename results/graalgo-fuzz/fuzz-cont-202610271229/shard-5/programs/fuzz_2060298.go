package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5)
        var v6 int = 128
        v7 := GenS1{F0: 16, F1: ("foo bar" + (("\n\t" + "foo bar") + "hello")), F2: 299}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 1
    _ = v8
    return (("" + "foo bar") >= "\n\t")
}
func genFn1() int64 {
    var v9 float32 = (((0.0 + 3.14159) + 0.0) + 0.0)
    var v10 int64 = 830
    switch ("\n\t" + (("a" + ("\n\t" + (("hello" + "hello") + "a"))) + "foo bar")) {
    case "a":
        v11 := GenS1{F0: 127, F1: ("\n\t" + (("foo bar" + "a") + "foo bar")), F2: 4294967295}
        _ = v11
    case "hello":
        switch ("foo bar" + "hello") {
        case "\n\t":
            for fi12 := 0; fi12 < 9; fi12++ {
                v10 += (^(v10))
                var v13 float64 = ((float64(float32(0.0)) + 0.0) + (3.14159 - 3.14159))
                var v14 uint8 = 128
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    case "":
        var v15 float32 = v9
        _ = v15
    default:
        if true {
            v10 = (v10 / ((v10) | 1))
            v10 = v10
        }
    }
    _ = v10
    _ = v9
    return (v10 % (((v10 * v10)) | 1))
}

func main() {
    var v1 int = 32767
    var v2 uint64 = 0
    var v3 bool = ((v2 << ((255) & 63)) == v2)
    var v4 uint64 = (v2 | 312)
    var v16 int64 = -594
    if (v3 || false) {
        fmt.Println(v4, v1)
        v16 = (^(v16))
    }
    if true {
        var v17 int32 = -775
        v16 += (v16 & 0)
        _ = v17
    }
    if (false && v3) {
        fmt.Println(v1, v16)
        v18 := GenS2{F0: v2}
        var v19 float32 = -2.25
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 int = (fi20 ^ (v1 / ((-1) | 1)))
        for fi22 := 0; fi22 < 12; fi22++ {
            switch (("a" + ("\n\t" + "foo bar")) + "hello") {
            case "hello":
                fmt.Println(v1)
            case "\n\t":
                v23 := GenS1{F0: 1, F1: "\n\t", F2: 42}
                _ = v23
            }
            var v24 uint8 = 217
            var v25 uint64 = v4
            var v26 uint32 = 65535
            var v27 uint32 = 0
            var v28 uint64 = 2
            fmt.Println(v1, v3)
            _ = fi22
            _ = v24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi20
        _ = v21
    }
    switch v1 {
    case -1:
        fmt.Println(v4, v3, v1)
    case -9223372036854775808:
        var v29 uint = 18446744073709551615
        _ = v29
    case -426:
        v30 := []int{-1}
        _ = v30
    }
    v31 := []int{}
    switch int(v4) {
    case 127:
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint32 = 325
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 int16 = 0
        _ = fi34
        _ = v35
    }
    var v36 int64 = v16
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 uint16 = 65535
            var v40 int = fi37
            v41 := GenS1{F0: 0, F1: "hello", F2: 2}
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi37
    }
    if v3 {
        var v42 float64 = 3.14159
        _ = v42
    } else {
        if (42 <= -385) {
            v1 -= (v1 >> ((65535) & 63))
            fmt.Println(v16, v4)
            if v3 {
                var v43 int = int(v2)
                _ = v43
            } else {
                var v44 bool = (("\n\t" + "hello") != ("\n\t" + ("hello" + "foo bar")))
                switch ("\n\t" + ("hello" + ("hello" + (("" + "foo bar") + ("\n\t" + "hello"))))) {
                case "hello":
                    var v45 float64 = (0.1 + (1.5 - float64(1.5)))
                    _ = v45
                default:
                    v31 = append(v31, 127)
                }
                _ = len(v31)
                _ = v44
            }
        } else {
            fmt.Println(v1)
            fmt.Println(v3)
        }
        _ = len(v31)
    }
    v46 := GenS1{F0: -1, F1: "hello", F2: 359}
    v47 := GenS3{F0: (v2 & (^(v2))), F1: 32767, F2: 0}
    var v48 int8 = -76
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v48
}
