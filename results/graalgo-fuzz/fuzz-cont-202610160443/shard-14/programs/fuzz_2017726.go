package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64 }
func genFn0() string {
    v5 := []int{}
    var v6 int16 = 0
    var v7 float64 = (3.14159 / (0.1 + 3.14159))
    v8 := GenS1{F0: "hello", F1: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint64, p1 string, p2 bool) uint16 {
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{(-(fi9)), fi9, (fi9 / (((fi9 - 9223372036854775807)) | 1)), (fi9 ^ -2147483648)}
        var v11 int16 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    p1 = (p1 + p1)
    switch p1 {
    case "a":
        p2 = false
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := []int{633, 2147483647, 207, -73}
    var v2 uint = 4294967295
    var v3 int16 = 32767
    var v4 bool = false
    v1 = append(v1, 32768)
    var v12 float32 = (1e10 - 1e10)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint8 = 42
        if ((v2 / ((v2) | 1)) < v2) {
            v15 := GenS2{F0: 0}
            _ = len(v1)
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    v16 := GenS1{F0: "foo bar", F1: 4294967295}
    var v17 int8 = 42
    var v18 string = "foo bar"
    v2 = v2
    v19 := GenS2{F0: -936}
    var v20 float64 = (1.5 / 1.5)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (-3 < -118) {
            if v4 {
                switch -801 {
                case 9223372036854775807:
                    v21 := GenS2{F0: 1}
                    _ = v21
                case 4294967295:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v12 = v12
                        _ = fi22
                    }
                default:
                    var v23 uint = (v2 % (((v2 / (((^(v2))) | 1))) | 1))
                    _ = v23
                }
                fmt.Println(v20)
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := GenS1{F0: "", F1: 4294967295}
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        } else {
            switch v18 {
            case "hello":
                fmt.Println(v2, v17, v3)
            case "":
                fmt.Println(v20, v3)
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                for fi28 := 0; fi28 < 1; fi28++ {
                    if (!(174 <= 366)) {
                        var v29 bool = v4
                        if len(v1) > 0 {
                            v1[0] = (fi28 ^ fi27)
                        }
                        _ = v29
                    }
                    v1 = append(v1, fi28)
                    var v30 string = v18
                    _ = fi28
                    _ = v30
                }
                _ = fi27
            }
        }
    } else {
        var v31 float32 = float32((v20 * (v20 / v20)))
        var v32 int8 = (v17 << ((4294967295) & 63))
        _ = v31
        _ = v32
    }
    var v33 int = -1
    v3 = (v3 * (v3 % ((32767) | 1)))
    var v34 int16 = v3
    var v35 uint8 = 0
    for fi36 := 0; fi36 < 7; fi36++ {
        v3 += (v3 >> ((v2) & 63))
        _ = fi36
    }
    v35 = (v35 << (((v2 % ((v2) | 1))) & 63))
    switch (v33 + v33) {
    case 764:
        for fi37 := 0; fi37 < 2; fi37++ {
            v38 := GenS1{F0: "a", F1: 155}
            switch v18 {
            case "foo bar":
                v4 = true
            default:
                v4 = v4
            }
            _ = fi37
            _ = v38
        }
    case -2147483648:
        switch v18 {
        case "hello":
            v34 += (v34 & (v34 | (v34 % (((v34 | 534)) | 1))))
        }
    case 648:
        if true {
            var v39 float64 = v20
            if v4 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v40 := GenS1{F0: v18, F1: 938}
                v35 = v35
                _ = v40
            }
            _ = v39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
