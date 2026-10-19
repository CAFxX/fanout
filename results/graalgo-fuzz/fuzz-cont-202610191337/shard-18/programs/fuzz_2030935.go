package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0() int8 {
    v6 := []int{2147483647, 32767, -2147483648}
    for fi7 := 0; fi7 < 7; fi7++ {
        if true {
            switch ("hello" + (("\n\t" + "") + "foo bar")) {
            case "":
                var v8 string = "foo bar"
                _ = v8
            }
            switch "foo bar" {
            case "":
                var v9 int = fi7
                _ = v9
            case "\n\t":
                var v10 int = (fi7 << ((255) & 63))
                _ = v10
            case "foo bar":
                v11 := GenS1{F0: (42 <= 1), F1: 714, F2: 42}
                _ = v11
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        var v12 uint64 = 0
        _ = fi7
        _ = v12
    }
    _ = v6
    return -128
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 0
    v4 := make(map[uint32]int)
    var v5 uint = 0
    fmt.Println(v1, v3, v2)
    var v13 bool = true
    var v14 string = ""
    if (!(v13 && (!true))) {
        switch 9223372036854775807 {
        case 42:
            if ((false && v13) || (v1 == 3.14159)) {
                fmt.Println(v1, v14, v13)
                v15 := GenS2{F0: (128 == 0)}
                _ = v15
            } else {
                var v16 bool = false
                v17 := []int{127}
                v16 = true
                _ = v16
                _ = v17
            }
        case 9223372036854775807:
            var v18 uint16 = 706
            _ = v18
        case 2147483647:
            fmt.Println(v2, v14, v14)
        default:
            for fi19 := 0; fi19 < 1; fi19++ {
                switch 128 {
                case 42:
                    v20 := GenS1{F0: v13, F1: 120, F2: 65535}
                    _ = v20
                case -9223372036854775808:
                    var v21 int8 = 57
                    _ = v21
                default:
                    var v22 int16 = 1000
                    _ = v22
                }
                v23 := []int{fi19}
                fmt.Println(v5)
                _ = fi19
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint32 = 209
            if v13 {
                var v26 bool = true
                v13 = (!false)
                _ = v26
            }
            v27 := []int{fi24}
            _ = fi24
            _ = v25
            _ = v27
        }
        fmt.Println(v5, v1, v3)
    } else {
        if false {
            v2 = v2
            v28 := GenS3{F0: -128}
            _ = v28
        }
        var v29 float32 = v1
        if v13 {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 float64 = 1e10
                switch ("\n\t" + (v14 + v14)) {
                case "hello":
                    var v32 string = (v14 + (v14 + ""))
                    _ = v32
                default:
                    var v33 int8 = 1
                    _ = v33
                }
                v13 = (v1 != v1)
                _ = fi30
                _ = v31
            }
        } else {
            v5 -= 297
        }
        _ = v29
    }
    fmt.Println(v14, v3)
    switch 127 {
    case 564:
        if true {
            v34 := GenS1{F0: ((v5 & (v5 - 1)) > 799), F1: v5, F2: 130}
            _ = v34
        }
    }
    fmt.Println(v14, v14)
    v3 = v3
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 uint64 = (v2 * (v2 * v2))
            v38 := GenS4{F0: -120, F1: 42}
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    if false {
        v39 := GenS1{F0: (v5 >= (v5 ^ 42)), F1: (^(v5)), F2: 0}
        fmt.Println(v3, v2, v14)
        _ = v39
    } else {
        fmt.Println(v14, v5, v3)
    }
    if true {
        fmt.Println(v14, v2, v13)
    }
    if ("" == "") {
        var v40 int16 = -1
        v1 += v1
        if true {
            var v41 int16 = v40
            var v42 uint64 = v2
            _ = v41
            _ = v42
        }
        _ = v40
    }
    v5 = v5
    switch -1 {
    case 32767:
        var v43 uint8 = 207
        _ = v43
    }
    var v44 uint8 = 1
    v13 = false
    switch v14 {
    case "a":
        if false {
            v45 := GenS2{F0: ((v44 > (v44 * 182)) || v13)}
            _ = v45
        }
    case "foo bar":
        v2 = v2
    default:
        v1 -= (v1 - ((((float32((-2.25 / (0.0 - float64(1e10)))) / v1) / (v1 / 1.0)) / v1) - -2.25))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
