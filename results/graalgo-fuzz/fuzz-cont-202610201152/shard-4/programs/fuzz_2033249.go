package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    v3 := []uint32{296, 0, 65535, 2147483648}
    v4 := []uint32{2}
    v1 = (^(v1))
    v1 = 128
    v4 = append(v4, 42)
    v1 += v1
    fmt.Println(v1)
    v5 := GenS1{F0: "a", F1: -168}
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint64 = 991
            fmt.Println(v1)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi6
            _ = v7
        }
    } else {
        switch "\n\t" {
        case "a":
            var v8 uint = 996
            _ = v8
        case "hello":
            var v9 int8 = -1
            _ = v9
        case "foo bar":
            v2[730] = (!((2147483647 < 825) && false))
        default:
            fmt.Println(v1)
        }
        var v10 int = 9223372036854775807
        fmt.Println(v10, v1)
        _ = v10
    }
    if false {
        if (0 != 0) {
            var v11 string = "\n\t"
            fmt.Println(v11)
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = v11
        } else {
            var v12 int32 = -2147483648
            switch 42 {
            case -446:
                var v13 int64 = 1
                _ = v13
            }
            switch 4294967295 {
            case 964:
                v2[0] = (990 == 0)
            }
            _ = v12
        }
        var v14 uint32 = 2147483648
        for fi15 := 0; fi15 < 8; fi15++ {
            v4 = append(v4, v14)
            v16 := GenS1{F0: "hello", F1: (fi15 >> ((1) & 63))}
            var v17 int8 = 100
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    } else {
        v1 += (v1 / (((v1 >> ((0) & 63))) | 1))
        v1 = 0
    }
    if false {
        switch 2147483647 {
        case 127:
            var v18 int = 127
            _ = v18
        case 4294967295:
            var v19 bool = false
            _ = v19
        }
        fmt.Println(v1)
        var v20 string = ""
        _ = v20
    }
    if (true || false) {
        v1 -= v1
        for fi21 := 0; fi21 < 11; fi21++ {
            switch "a" {
            case "foo bar":
                var v22 int32 = -2147483648
                _ = v22
            case "":
                v1 = 242
            }
            switch (("a" + "a") + "foo bar") {
            case "":
                v1 += (v1 - v1)
            case "a":
                var v23 uint64 = 1
                _ = v23
            default:
                fmt.Println(fi21, v1)
            }
            _ = fi21
        }
    } else {
        if true {
            v24 := GenS2{F0: v1, F1: 1}
            v25 := GenS1{F0: "\n\t", F1: 4294967295}
            if (true && (false && true)) {
                v1 -= (v1 << ((18446744073709551615) & 63))
                var v26 uint16 = 0
                _ = v26
            }
            _ = v24
            _ = v25
        }
        var v27 uint16 = 207
        _ = v27
    }
    if (1 >= -9223372036854775808) {
        fmt.Println(v1)
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v1, fi28)
            var v29 float32 = 1.5
            v29 -= v29
            _ = fi28
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 6; fi30++ {
            v31 := GenS3{F0: 100, F1: fi30}
            _ = fi30
            _ = v31
        }
        var v32 uint8 = (v1 ^ 154)
        _ = v32
    }
    if (false && true) {
        v33 := GenS4{F0: 42, F1: 791}
        v34 := GenS2{F0: (v1 / (((v1 & v1)) | 1)), F1: -701}
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int8 = 100
        var v37 int16 = 32767
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int8 = -112
    v1 -= (v1 << ((0) & 63))
    v39 := GenS2{F0: (v1 << ((208) & 63)), F1: 824}
    fmt.Println(v38)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
