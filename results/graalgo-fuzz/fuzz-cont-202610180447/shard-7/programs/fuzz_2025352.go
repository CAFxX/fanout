package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0() bool {
    var v7 uint16 = 32767
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(fi8, v7)
        fmt.Println(fi8, fi8)
        var v9 int = (fi8 | (fi8 ^ (fi8 / ((fi8) | 1))))
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    _ = v7
    return false
}

func main() {
    var v1 uint = 255
    var v2 int = 4294967295
    v3 := make(map[int]string)
    var v4 uint64 = 23
    v5 := make(map[string]int)
    var v6 int = 0
    v6 = -75
    fmt.Println(v1)
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 42, F1: "hello", F2: 504}
        var v12 string = "hello"
        v2 = int(v4)
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v6)
    if (!true) {
        if (true || false) {
            v13 := GenS2{F0: false}
            v1 += 0
            _ = v13
        } else {
            v14 := GenS2{F0: (false || false)}
            _ = v14
        }
        v15 := GenS2{F0: (42 <= -128)}
        switch ("a" + "\n\t") {
        case "\n\t":
            _ = len(v3)
        }
        _ = v15
    }
    v16 := GenS3{F0: 1, F1: 65535, F2: 1}
    v6 -= 325
    if (3.14159 <= (3.14159 + 3.14159)) {
        if true {
            var v17 string = ("foo bar" + "\n\t")
            var v18 float32 = 1e10
            switch v17 {
            case "a":
                var v19 float32 = (0.0 / 1.0)
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 float64 = 1.5
                    var v22 uint = v1
                    v2 = (-(v2))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case "foo bar":
                var v23 uint = v1
                _ = v23
            default:
                _ = len(v5)
            }
            _ = v17
            _ = v18
        } else {
            fmt.Println(v4, v1, v2)
            var v24 string = "foo bar"
            _ = v24
        }
        v25 := []int{32768, -125, (v6 + 32767)}
        _ = v25
    } else {
        var v26 uint32 = 4294967295
        _ = len(v3)
        delete(v3, 42)
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := v5["hello"]
            v29 := GenS3{F0: -1, F1: 4294967295, F2: 128}
            _ = fi27
            _ = v28
            _ = v29
        }
        fmt.Println(v2)
    }
    fmt.Println(v4, v1)
    var v30 string = "a"
    for fi31 := 0; fi31 < 19; fi31++ {
        if (!((^(v1)) != 2)) {
            var v32 bool = false
            _ = v32
        }
        var v33 bool = false
        for fi34 := 0; fi34 < 7; fi34++ {
            switch (v30 + (v30 + v30)) {
            case "foo bar":
                fmt.Println(v6, v2, v1)
            default:
                switch v30 {
                case "hello":
                    v35 := []int{-2147483648, v2, 1, 0}
                    _ = v35
                default:
                    var v36 int64 = 42
                    _ = v36
                }
            }
            var v37 uint = (v1 >> ((v1) & 63))
            var v38 uint = v37
            var v39 string = v30
            _ = fi34
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi31
        _ = v33
    }
    fmt.Println(v2, v1)
    if false {
        fmt.Println(v6, v30)
        if (true && (2147483647 != 746)) {
            var v40 int8 = -19
            _ = v40
        } else {
            v2 = (-(v2))
            fmt.Println(v4)
            for fi41 := 0; fi41 < 7; fi41++ {
                switch v30 {
                case "foo bar":
                    v6 = -9223372036854775808
                case "hello":
                    fmt.Println(fi41, v6, fi41)
                case "a":
                    v42 := GenS4{F0: 2}
                    _ = v42
                }
                _ = fi41
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
