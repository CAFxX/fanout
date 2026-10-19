package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 float64 = 1.5
    var v2 uint = 255
    var v3 int64 = 42
    if false {
        v3 -= v3
    } else {
        fmt.Println(v1, v1, v2)
        switch 984 {
        case -1:
            switch 42 {
            case 2147483647:
                v4 := []int{}
                _ = v4
            }
        case -2147483648:
            v5 := GenS1{F0: "hello"}
            _ = v5
        case -984:
            switch 4294967295 {
            case -224:
                var v6 int = 9223372036854775807
                _ = v6
            default:
                switch 2147483647 {
                case 9223372036854775807:
                    if ((false || (false || (32767 < 32767))) && true) {
                        v7 := []int{}
                        v8 := GenS1{F0: ""}
                        _ = v7
                        _ = v8
                    }
                case 391:
                    v2 -= (v2 & 398)
                case 0:
                    for fi9 := 0; fi9 < 12; fi9++ {
                        fmt.Println(v3, v2, fi9)
                        v10 := GenS1{F0: "foo bar"}
                        v1 = (((1e10 - v1) * v1) / 3.14159)
                        _ = fi9
                        _ = v10
                    }
                }
            }
        default:
            var v11 float32 = 0.1
            _ = v11
        }
        if ((!true) && false) {
            v12 := []int{-360, 128, -822}
            v13 := GenS2{F0: ""}
            v2 = 521
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v2, v2, v1)
    if (!(v3 >= v3)) {
        var v14 int64 = (v3 - 9223372036854775807)
        for fi15 := 0; fi15 < 18; fi15++ {
            if false {
                v2 = (v2 & v2)
            }
            v14 = int64(v2)
            if false {
                for fi16 := 0; fi16 < 14; fi16++ {
                    v14 -= -574
                    v17 := GenS2{F0: "\n\t"}
                    if (2 >= v2) {
                        v18 := GenS3{F0: 127, F1: 1, F2: 720}
                        switch (-(fi15)) {
                        case 9223372036854775807:
                            var v19 bool = (!true)
                            _ = v19
                        case -9223372036854775808:
                            if false {
                                var v20 uint = v2
                                var v21 int32 = 1
                                v20 -= v2
                                _ = v20
                                _ = v21
                            }
                        }
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
        switch 127 {
        case 32767:
            v22 := []int{42, 127, 127, 9223372036854775807}
            _ = v22
        default:
            var v23 uint32 = 2147483648
            _ = v23
        }
        _ = v14
    } else {
        v1 -= 1e10
        v24 := []int{-1, -1, 32767}
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int32 = 1
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if true {
            v1 = v1
            var v28 uint = 65535
            _ = v28
        } else {
            v29 := []int{}
            _ = v29
        }
        _ = fi27
    }
    var v30 int64 = -1
    v31 := []int{4294967295, 128, 1}
    v1 -= float64(0.0)
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
