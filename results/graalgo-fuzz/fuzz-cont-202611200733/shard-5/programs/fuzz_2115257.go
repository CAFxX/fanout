package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }

func main() {
    var v1 int32 = 404
    v2 := []int{-281}
    v3 := make(map[uint32]string)
    var v4 uint32 = 627
    v5 := []string{"foo bar"}
    var v6 int8 = 1
    var v7 int8 = (v6 ^ (-(v6)))
    fmt.Println(v1, v4, v4)
    var v8 uint16 = 0
    _ = len(v2)
    if false {
        switch ("\n\t" + "a") {
        case "":
            var v9 bool = false
            _ = v9
        }
        v10 := GenS1{F0: -880, F1: "hello"}
        var v11 uint32 = (v4 >> ((2) & 63))
        _ = v10
        _ = v11
    }
    switch ("a" + "a") {
    case "foo bar":
        var v12 int32 = v1
        _ = v12
    default:
        fmt.Println(v4, v4, v7)
    }
    _ = len(v2)
    v1 -= -1
    var v13 int64 = 335
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 bool = ((v13 >> ((4294967295) & 63)) == (v13 << ((605) & 63)))
        for fi16 := 0; fi16 < 8; fi16++ {
            if (!(539 >= 42)) {
                v17 := GenS1{F0: 42, F1: "\n\t"}
                _ = v17
            }
            var v18 float64 = (1e10 * (1.5 / 3.14159))
            _ = fi16
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (("\n\t" > (("a" + "\n\t") + ("\n\t" + "foo bar"))) && (!false)) {
                if v15 {
                    if v15 {
                        var v20 bool = true
                        v15 = (true || (!false))
                        var v21 int64 = v13
                        _ = len(v2)
                        var v22 int16 = -1
                        v4 += (^(v4))
                        var v23 uint64 = 1
                        if (!(!(v23 > (v23 * (v23 / ((477) | 1)))))) {
                            var v24 bool = (v7 < -1)
                            var v25 uint64 = 949
                            v26 := GenS1{F0: v22, F1: "\n\t"}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    switch (fi14 * -1) {
                    case -832:
                        var v27 uint64 = 18446744073709551615
                        _ = v27
                    case 2147483647:
                        v28 := GenS1{F0: -883, F1: "\n\t"}
                        _ = v28
                    default:
                        var v29 uint32 = v4
                        v6 -= v6
                        _ = v29
                    }
                    v30 := GenS2{F0: v7}
                    _ = v30
                } else {
                    var v31 int32 = -2147483648
                    var v32 uint32 = (v4 ^ (v4 % ((65535) | 1)))
                    var v33 uint = 42
                    _ = v31
                    _ = v32
                    _ = v33
                }
                v3[v4] = "a"
                switch (-(fi19)) {
                case -857:
                    v4 += uint32(v8)
                }
            } else {
                switch fi14 {
                case 4294967295:
                    v1 += (v1 ^ v1)
                }
                v13 = (v13 & (v13 + (-(v13))))
            }
            _ = fi19
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    if true {
        var v34 int8 = int8(v13)
        v2 = append(v2, 127)
        _ = v34
    }
    var v35 bool = true
    v36 := GenS3{F0: v35, F1: 0, F2: -1}
    v8 = v8
    if false {
        fmt.Println(v13)
    } else {
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        fmt.Println(v8)
        v7 -= 42
    }
    switch "" {
    case "foo bar":
        var v37 float64 = ((0.1 - (1.5 - 1.5)) / -2.25)
        _ = v37
    case "hello":
        var v38 string = "a"
        _ = v38
    case "\n\t":
        var v39 uint64 = 0
        _ = v39
    }
    v6 = (v7 + (v7 - (^(v7))))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
