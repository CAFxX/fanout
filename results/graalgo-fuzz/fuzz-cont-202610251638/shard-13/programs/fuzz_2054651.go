package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint = 1
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    v6 := []int{2147483647, 2147483647, 127, 774}
    _ = len(v6)
    var v7 uint = v3
    v8 := GenS1{F0: (("foo bar" + (((("" + "hello") + "foo bar") + "foo bar") + "foo bar")) + ("\n\t" + ("a" + ("" + ("hello" + ""))))), F1: false, F2: 2147483647}
    var v9 uint16 = 2
    var v10 uint64 = 18446744073709551615
    var v11 bool = (v7 > v7)
    v12 := GenS1{F0: "\n\t", F1: false, F2: -1}
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v3, v7, v3)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi13
    }
    if v11 {
        for fi14 := 0; fi14 < 20; fi14++ {
            v10 -= 2
            v2[""] = fi14
            _ = fi14
        }
    }
    var v15 int64 = 9223372036854775807
    if (v11 && ((-410 >= 42) || (0.1 >= 1.5))) {
        fmt.Println(v15, v7)
        v11 = (v11 && v11)
    }
    fmt.Println(v15)
    v9 = 42
    if v11 {
        switch "foo bar" {
        case "foo bar":
            for fi16 := 0; fi16 < 3; fi16++ {
                delete(v2, "a")
                _ = fi16
            }
        case "a":
            switch ("foo bar" + "hello") {
            case "":
                delete(v5, 2)
            case "hello":
                var v17 uint8 = 1
                _ = v17
            case "foo bar":
                var v18 int8 = 88
                _ = v18
            default:
                var v19 float32 = (1e10 / 1.0)
                _ = v19
            }
        default:
            if v11 {
                v15 -= (v15 * 0)
            } else {
                v20 := GenS1{F0: ("hello" + "\n\t"), F1: false, F2: 0}
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            switch "\n\t" {
            case "foo bar":
                v7 = (v3 >> ((v3) & 63))
            case "\n\t":
                fmt.Println(v3, v4)
            }
            var v22 int8 = 100
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: 132, F1: (v10 % (((v10 / ((v10) | 1))) | 1)), F2: (v9 >> ((v3) & 63))}
        _ = v23
    }
    var v24 uint32 = 1
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := v2[("foo bar" + "hello")]
        var v27 int16 = 1000
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 uint8 = 42
            var v30 bool = (v27 == (v27 % ((int16(v15)) | 1)))
            v31 := GenS3{F0: ("\n\t" + "hello"), F1: 42}
            v32 := GenS1{F0: ("a" + "\n\t"), F1: ((v7 % (((v7 - 0)) | 1)) == v7), F2: fi28}
            var v33 float32 = 0.1
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    if v11 {
        fmt.Println(v4)
        var v34 bool = v11
        v15 += (v15 * (v15 & (v15 >> ((v3) & 63))))
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
