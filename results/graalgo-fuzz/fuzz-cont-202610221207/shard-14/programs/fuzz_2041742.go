package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 string; F1 uint32 }
func genFn0(p0 uint64) int {
    var v5 float32 = -2.25
    v5 = v5
    var v6 uint32 = 4294967295
    v6 -= (v6 << ((1) & 63))
    var v7 int = -1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (-(v7))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    var v3 string = ("\n\t" + "\n\t")
    var v4 uint8 = 2
    if true {
        var v8 string = ("" + v3)
        if (!(-9223372036854775808 == genFn0(2))) {
            v2 -= (-(v2))
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int16 = -415
                _ = fi9
                _ = v10
            }
        } else {
            v2 = (v2 ^ (^(v2)))
        }
        _ = v8
    } else {
        if ((v3 + "") <= "\n\t") {
            v11 := GenS1{F0: (42 == -128), F1: 42}
            if (-32768 < 32767) {
                fmt.Println(v4, v3)
            } else {
                var v12 float32 = ((1e10 - 3.14159) + 1.5)
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint16 = 32767
                    v4 = 128
                    v15 := GenS1{F0: false, F1: -43}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := []int{}
                _ = v12
                _ = v16
            }
            var v17 float32 = float32(0.1)
            _ = v11
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                v2 = 42
                v19 := GenS2{F0: 1}
                _ = fi18
                _ = v19
            }
            v20 := GenS3{F0: 72}
            _ = v20
        }
        v21 := v1[154]
        fmt.Println(v3, v4)
        _ = v21
    }
    v22 := []int{}
    v23 := GenS3{F0: 37}
    var v24 int32 = -913
    fmt.Println(v4)
    if len(v22) > 0 {
        v22[0] = 47
    }
    switch "hello" {
    case "\n\t":
        for fi25 := 0; fi25 < 19; fi25++ {
            if (!true) {
                switch "foo bar" {
                case "":
                    var v26 int32 = 787
                    _ = v26
                }
                var v27 int16 = -32768
                _ = v27
            } else {
                if true {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        v29 := GenS1{F0: (v4 != (v4 << ((255) & 63))), F1: -56}
                        _ = fi28
                        _ = v29
                    }
                }
                if len(v22) > 0 {
                    v22[0] = (fi25 * fi25)
                }
            }
            _ = len(v22)
            _ = fi25
        }
    case "foo bar":
        if ((false || (-9223372036854775808 < (v2 % (((-(v2))) | 1)))) && true) {
            delete(v1, 1)
        }
    default:
        for fi30 := 0; fi30 < 2; fi30++ {
            switch -1 {
            case 127:
                fmt.Println(v24, v4)
            case 128:
                v31 := GenS1{F0: (!false), F1: 1}
                _ = v31
            case 823:
                fmt.Println(v24, v3)
            default:
                var v32 string = v3
                _ = v32
            }
            if len(v22) > 0 {
                v22[0] = (-(fi30))
            }
            var v33 uint64 = 18446744073709551615
            _ = fi30
            _ = v33
        }
    }
    v34 := GenS1{F0: true, F1: -128}
    fmt.Println(v2, v24, v2)
    var v35 uint8 = v4
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 5; fi37++ {
            switch (fi36 * (fi37 >> ((169) & 63))) {
            case 4294967295:
                v38 := GenS4{F0: 65535, F1: (v4 / ((v4) | 1))}
                _ = v38
            case 2147483647:
                switch (fi37 - (^(fi36))) {
                case -824:
                    v39 := GenS4{F0: 2, F1: (v4 ^ v35)}
                    _ = v39
                }
            }
            if false {
                var v40 string = (v3 + v3)
                var v41 int = (fi36 ^ (^(fi37)))
                v42 := v1[4294967295]
                _ = v40
                _ = v41
                _ = v42
            }
            if true {
                v43 := GenS5{F0: "a", F1: uint32(v35)}
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = v43
            }
            var v44 int8 = 42
            _ = fi37
            _ = v44
        }
        _ = fi36
    }
    fmt.Println(v35, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
