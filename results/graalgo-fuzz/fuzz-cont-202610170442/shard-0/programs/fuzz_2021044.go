package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch 289 {
        case 2147483647:
            for fi5 := 0; fi5 < 3; fi5++ {
                v6 := GenS1{F0: 1, F1: 32767, F2: -2147483648}
                _ = fi5
                _ = v6
            }
        default:
            var v7 uint16 = 0
            _ = v7
        }
        fmt.Println(fi4)
        var v8 float32 = 0.1
        _ = fi4
        _ = v8
    }
    var v9 bool = (!true)
    switch -9223372036854775808 {
    case -158:
        v10 := []int{99, 154, -2147483648}
        _ = v10
    case -9223372036854775808:
        var v11 string = ""
        _ = v11
    case 128:
        fmt.Println(v9)
    default:
        switch "a" {
        case "":
            v9 = true
        case "a":
            if v9 {
                v9 = false
            }
        case "hello":
            var v12 int8 = 127
            _ = v12
        default:
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int8 = 42
                v15 := []int{(fi13 % (((fi13 << ((65535) & 63))) | 1)), (fi13 + (fi13 ^ (-(fi13)))), fi13}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 uint32 = 81
    v17 := GenS2{F0: 2, F1: ("" + "hello"), F2: 966}
    _ = v16
    _ = v17
    _ = v9
    return 1
}
func genFn1() int {
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 float64 = (-2.25 + 1e10)
        fmt.Println(fi18)
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 65535
    v20 -= v20
    var v21 float32 = ((0.1 + 1.5) / 0.1)
    _ = v20
    _ = v21
    return 837
}

func main() {
    var v1 string = ("a" + "a")
    v2 := make(map[int]string)
    var v3 int = -181
    var v22 string = v1
    fmt.Println(v3, v22)
    var v23 int = (v3 >> ((42) & 63))
    _ = len(v2)
    v22 = ((("foo bar" + v1) + v1) + v22)
    v23 = v3
    for fi24 := 0; fi24 < 1; fi24++ {
        switch v1 {
        case "foo bar":
            v22 = ""
        case "hello":
            v25 := []int{fi24}
            _ = v25
        default:
            v26 := []int{(v3 - (v3 / (((v23 - v23)) | 1))), (-(v3))}
            _ = v26
        }
        var v27 uint8 = 9
        for fi28 := 0; fi28 < 10; fi28++ {
            switch v23 {
            case 32767:
                var v29 uint16 = 2
                delete(v2, fi28)
                var v30 uint16 = (^(v29))
                _ = v29
                _ = v30
            }
            var v31 uint8 = (v27 & (v27 + v27))
            var v32 int = genFn1()
            var v33 int16 = 1
            v33 += (v33 ^ (-(v33)))
            if (v1 == (v1 + v1)) {
                var v34 int64 = 0
                v35 := GenS3{F0: 546, F1: v27, F2: v31}
                v36 := []int{v3, -9223372036854775808}
                _ = v34
                _ = v35
                _ = v36
            } else {
                switch "foo bar" {
                case "hello":
                    v37 := GenS4{F0: 9223372036854775807}
                    v27 = v31
                    _ = v37
                case "foo bar":
                    v38 := GenS5{F0: 1, F1: v27}
                    _ = v38
                case "":
                    var v39 int64 = -9223372036854775808
                    _ = v39
                }
            }
            _ = fi28
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi24
        _ = v27
    }
    var v40 int64 = 42
    if true {
        switch (v23 | (v3 >> ((672) & 63))) {
        case 449:
            var v41 int8 = 42
            _ = v41
        case -1:
            var v42 float32 = 1e10
            _ = v42
        case -9223372036854775808:
            var v43 int16 = 32767
            _ = v43
        }
        v44 := []int{128}
        _ = v44
    } else {
        v45 := []int{(v3 | v23), (v3 - v23)}
        _ = v45
    }
    var v46 uint32 = 65535
    v40 += (v40 - v40)
    fmt.Println(v1)
    v3 = v23
    if true {
        v3 += v23
        var v47 float32 = float32(((0.0 / 1.0) - -2.25))
        fmt.Println(v46, v1)
        _ = v47
    } else {
        if (32767 != genFn0()) {
            v48 := []int{v3, (v23 & 9223372036854775807)}
            switch v3 {
            case 830:
                v3 += -1
            case 32768:
                v46 += 120
            case 348:
                v1 = "foo bar"
            }
            v3 += v3
            _ = v48
        } else {
            switch "\n\t" {
            case "foo bar":
                switch "foo bar" {
                case "foo bar":
                    v3 += (v23 / (((-(v23))) | 1))
                case "":
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 bool = false
                        v23 -= (-(fi49))
                        _ = fi49
                        _ = v50
                    }
                }
            case "":
                v46 -= v46
            default:
                fmt.Println(v23, v46, v40)
            }
            if false {
                fmt.Println(v40, v46)
            } else {
                if false {
                    v51 := GenS4{F0: (v40 + v40)}
                    v3 = int(v40)
                    _ = v51
                }
            }
            v2[int(v46)] = ""
        }
    }
    v52 := GenS5{F0: 0, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v40
    _ = v46
    _ = v52
}
