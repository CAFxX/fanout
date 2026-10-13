package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 int, p1 bool) int {
    var v6 uint = 2
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 0, F1: (9223372036854775807 == -512), F2: 32767}
        var v10 int8 = -128
        v10 = 126
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 uint8 = 42
    v11 := GenS2{F0: -76}
    var v12 int16 = 32767
    var v13 float64 = 0.1
    if false {
        if false {
            v14 := GenS2{F0: 127}
            v12 = -103
            if false {
                v13 = v13
                v12 -= 1000
            } else {
                for fi15 := 0; fi15 < 17; fi15++ {
                    if false {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v17 := v4[""]
                            v18 := GenS1{F0: 102, F1: (!true), F2: 32767}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v5 = 128
                    }
                    switch fi15 {
                    case 127:
                        v12 -= v12
                    case -2147483648:
                        switch (fi15 & (fi15 >> ((4294967295) & 63))) {
                        case -1:
                            var v19 string = ("a" + "hello")
                            _ = v19
                        }
                    case 1:
                        var v20 float32 = 3.14159
                        _ = v20
                    }
                    var v21 int32 = 42
                    _ = fi15
                    _ = v21
                }
                switch 128 {
                case 128:
                    switch "a" {
                    case "foo bar":
                        var v22 uint64 = 2
                        _ = v22
                    default:
                        v23 := GenS1{F0: 83, F1: (2147483647 != 1), F2: 716}
                        _ = v23
                    }
                case -2147483648:
                    delete(v3, ("\n\t" + "foo bar"))
                case 43:
                    fmt.Println(v13)
                default:
                    switch "\n\t" {
                    case "hello":
                        var v24 int32 = 2147483647
                        _ = v24
                    case "foo bar":
                        if false {
                            v5 = uint8(v1)
                        } else {
                            v25 := GenS3{F0: 42}
                            _ = v25
                        }
                    case "\n\t":
                        var v26 float32 = 0.0
                        _ = v26
                    }
                }
            }
            _ = v14
        } else {
            v12 -= v12
        }
        fmt.Println(v5, v13)
    }
    v5 -= 128
    _ = len(v4)
    if (-2147483648 != -568) {
        v1 = v1
        fmt.Println(v12)
    }
    v1 -= (^(v1))
    fmt.Println(v1)
    switch -1 {
    case 824:
        switch ("\n\t" + "hello") {
        case "":
            switch 171 {
            case -837:
                v27 := []int{381, 0, -266}
                _ = v27
            default:
                v1 -= (v1 << ((v1) & 63))
            }
        case "a":
            switch "hello" {
            case "":
                v28 := GenS3{F0: int(v5)}
                _ = v28
            case "foo bar":
                v29 := GenS2{F0: 127}
                _ = v29
            }
        case "hello":
            v12 = v12
        default:
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint = (v1 % ((v1) | 1))
                switch "foo bar" {
                case "a":
                    v32 := []int{fi30, genFn0(genFn0((^(fi30)), false), true)}
                    _ = v32
                case "\n\t":
                    var v33 uint16 = 2
                    _ = v33
                }
                if (true && (v31 > v1)) {
                    v5 = (v5 >> ((v31) & 63))
                } else {
                    var v34 int = (fi30 / ((1) | 1))
                    _ = v34
                }
                _ = fi30
                _ = v31
            }
        }
    case 791:
        v35 := GenS2{F0: -1}
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 4; fi37++ {
            v5 += v5
            v38 := GenS2{F0: 57}
            _ = fi37
            _ = v38
        }
        delete(v2, fi36)
        var v39 int8 = -1
        _ = fi36
        _ = v39
    }
    fmt.Println(v12, v5, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
