package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 uint }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 uint64 = 2
    var v5 float32 = 1e10
    var v6 uint32 = 323
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        for fi7 := 0; fi7 < 18; fi7++ {
            v5 = (((float32(3.14159) - (v5 + (v5 / v5))) / (v5 + float32(1e10))) * (v5 / 1e10))
            var v8 int32 = -1
            _ = fi7
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 15; fi9++ {
            v3["a"] = ("foo bar" + "a")
            _ = fi9
        }
    default:
        switch "foo bar" {
        case "hello":
            v10 := []int{}
            _ = v10
        case "":
            if false {
                v4 = v4
            }
        default:
            v11 := GenS1{F0: 475}
            _ = v11
        }
    }
    v6 -= v6
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS2{F0: (v6 | (v6 >> ((18446744073709551615) & 63))), F1: 347}
            v6 = (v6 >> ((2) & 63))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint8 = 159
        v5 = 1e10
        if true {
            if false {
                switch (fi15 & (fi15 * (fi15 >> ((989) & 63)))) {
                case 32767:
                    v17 := GenS3{F0: (fi15 / (((fi15 | (fi15 % ((128) | 1)))) | 1)), F1: 759, F2: 106}
                    _ = v17
                case 2147483647:
                    v18 := GenS4{F0: -1, F1: 1}
                    _ = v18
                }
                var v19 uint = 42
                _ = v19
            } else {
                v20 := GenS4{F0: 1, F1: 801}
                v21 := GenS2{F0: v6, F1: 73}
                if true {
                    var v22 uint16 = 215
                    _ = v22
                }
                _ = v20
                _ = v21
            }
            var v23 int8 = -128
            _ = v23
        }
        _ = fi15
        _ = v16
    }
    v4 = (v4 + v4)
    _ = len(v1)
    v24 := GenS3{F0: 0, F1: -9223372036854775808, F2: 42}
    switch 331 {
    case 127:
        if (((float64((v5 * (v5 + 0.1))) <= 1e10) || true) || (1 >= 0)) {
            var v25 uint64 = (^(v4))
            for fi26 := 0; fi26 < 14; fi26++ {
                switch "\n\t" {
                case "foo bar":
                    v4 += (v25 * v4)
                }
                fmt.Println(v6, v5, v4)
                _ = fi26
            }
            _ = v25
        }
    case 42:
        v27 := GenS1{F0: -9223372036854775808}
        _ = v27
    case 32767:
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := []int{}
            var v30 bool = (v4 < 2)
            v31 := GenS4{F0: -271, F1: 468}
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 5; fi32++ {
            switch "hello" {
            case "\n\t":
                var v33 uint32 = v6
                _ = v33
            case "a":
                for fi34 := 0; fi34 < 20; fi34++ {
                    v35 := GenS2{F0: v6, F1: 255}
                    v36 := []int{(fi34 - fi32), (fi34 - int(v6))}
                    if (!((42 >= 127) && (false && (751 >= 65535)))) {
                        var v37 uint8 = 2
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "hello":
                v38 := GenS1{F0: fi32}
                _ = v38
            }
            var v39 uint = 1
            _ = fi32
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        if false {
            for fi41 := 0; fi41 < 18; fi41++ {
                v42 := GenS2{F0: uint32(v4), F1: 42}
                _ = fi41
                _ = v42
            }
        }
        v43 := []int{(fi40 - fi40)}
        _ = fi40
        _ = v43
    }
    if (false && false) {
        if true {
            v5 += 1e10
            var v44 bool = true
            _ = v44
        }
        var v45 string = "foo bar"
        _ = v45
    }
    fmt.Println(v6, v5, v5)
    var v46 string = ""
    fmt.Println(v46, v46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
