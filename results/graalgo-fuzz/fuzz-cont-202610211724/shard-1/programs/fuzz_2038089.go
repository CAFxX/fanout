package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) uint8 {
    var v4 int = -976
    v5 := []int{v4, v4}
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = -128
    var v3 uint64 = 904
    v3 += (v3 << ((1) & 63))
    var v6 uint64 = v3
    switch "a" {
    case "":
        v7 := GenS1{F0: "a"}
        _ = v7
    case "foo bar":
        switch "" {
        case "\n\t":
            if (-9223372036854775808 < 554) {
                for fi8 := 0; fi8 < 6; fi8++ {
                    if true {
                        v6 = (v6 & v6)
                    } else {
                        fmt.Println(fi8)
                        v9 := []int{1}
                        var v10 uint64 = v3
                        _ = v10
                        _ = v9
                    }
                    var v11 int64 = 135
                    fmt.Println(fi8)
                    _ = fi8
                    _ = v11
                }
            }
        default:
            v12 := GenS1{F0: "\n\t"}
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 11; fi13++ {
            if true {
                var v14 uint = 44
                var v15 uint8 = genFn0(42)
                var v16 int8 = v2
                _ = v14
                _ = v15
                _ = v16
            } else {
                switch "" {
                case "foo bar":
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := GenS1{F0: "a"}
                        _ = fi17
                        _ = v18
                    }
                default:
                    v19 := []int{(fi13 >> ((18446744073709551615) & 63)), fi13}
                    _ = v19
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    v1 += (v1 & (v1 / ((uint32(v3)) | 1)))
                    if true {
                        fmt.Println(v1, fi13, v1)
                        v21 := GenS1{F0: (("a" + "a") + "")}
                        fmt.Println(v6, v3, fi20)
                        v2 = (v2 & v2)
                        fmt.Println(fi20, v6, v6)
                        if (true || (!(!(false || (v1 == v1))))) {
                            v6 -= (v3 ^ (^(v6)))
                            v3 -= (v6 >> ((463) & 63))
                            v6 = 0
                            v22 := []int{(fi20 % ((fi13) | 1)), (fi13 & fi13), fi13}
                            _ = v22
                        }
                        _ = v21
                    } else {
                        v23 := GenS1{F0: ("hello" + ("a" + "foo bar"))}
                        v6 = (^(v3))
                        _ = v23
                    }
                    var v24 int8 = 42
                    _ = fi20
                    _ = v24
                }
            }
            var v25 bool = true
            _ = fi13
            _ = v25
        }
    }
    switch 9223372036854775807 {
    case 0:
        var v26 int64 = 0
        _ = v26
    case 56:
        if true {
            v27 := GenS1{F0: "\n\t"}
            v28 := []int{944}
            _ = v27
            _ = v28
        } else {
            v29 := []int{127, 127}
            v3 += 42
            var v30 int64 = 531
            _ = v29
            _ = v30
        }
    }
    if false {
        var v31 int32 = 2147483647
        v1 = 1
        _ = v31
    } else {
        var v32 uint = 65535
        v2 -= v2
        v33 := []int{774, -9223372036854775808, 1}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v6 = v3
        var v35 int32 = -2147483648
        _ = fi34
        _ = v35
    }
    v36 := GenS2{F0: -1}
    v37 := []int{}
    v1 = (v1 >> ((2) & 63))
    if false {
        var v38 float32 = -2.25
        v37 = append(v37, -773)
        _ = v38
    } else {
        _ = len(v37)
    }
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v6
}
