package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{858, 961, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 941
    v4 := []int{444}
    v5 := make(map[int]string)
    switch "\n\t" {
    case "a":
        v3 += v3
    }
    var v6 bool = (1.5 != 3.14159)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: ("" + "")}
        _ = fi7
        _ = v8
    }
    if (42 == 42) {
        fmt.Println(v6, v3)
        var v9 uint8 = 2
        v3 = 0
        _ = v9
    }
    v6 = v6
    if v6 {
        var v10 string = "hello"
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = v10
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint = 255
            if v6 {
                v6 = v6
                if false {
                    if (!((!false) && v6)) {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            v3 = v3
                            v14 := GenS2{F0: -32768, F1: 128}
                            v15 := GenS1{F0: (("\n\t" + ("a" + "hello")) + "\n\t")}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                }
                v12 = (v12 - v12)
            } else {
                v16 := GenS1{F0: (("foo bar" + ("foo bar" + "foo bar")) + ("\n\t" + "\n\t"))}
                var v17 int16 = -776
                _ = len(v4)
                _ = v16
                _ = v17
            }
            v12 += 339
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v3, v3)
    v18 := GenS2{F0: -731, F1: 22}
    switch "foo bar" {
    case "":
        var v19 int = 2147483647
        _ = v19
    default:
        v20 := GenS2{F0: 1000, F1: 1}
        _ = v20
    }
    v6 = false
    if ((!false) || ((!((!v6) || v6)) && (42 < 393))) {
        v3 -= v3
        v6 = v6
    } else {
        if true {
            if (true && v6) {
                fmt.Println(v3, v6)
                v3 += (^(v3))
            }
            switch "foo bar" {
            case "\n\t":
                v21 := GenS2{F0: 1000, F1: 255}
                _ = v21
            default:
                for fi22 := 0; fi22 < 3; fi22++ {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        switch "\n\t" {
                        case "hello":
                            v3 += (v3 | (^(v3)))
                        case "":
                            fmt.Println(v6)
                        case "\n\t":
                            if v6 {
                                fmt.Println(v6)
                            } else {
                                var v24 float32 = 0.1
                                _ = v24
                            }
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v3 = v3
                            }
                        }
                        var v25 float64 = (1.5 / float64(3.14159))
                        _ = fi23
                        _ = v25
                    }
                    var v26 int16 = -1
                    delete(v2, 2147483648)
                    _ = fi22
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 1; fi27++ {
                delete(v2, 2)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi27
            }
        }
        var v28 int8 = 42
        _ = v28
    }
    switch 9223372036854775807 {
    case 32767:
        var v29 int16 = 1000
        _ = v29
    }
    v30 := GenS1{F0: "a"}
    var v31 bool = v6
    v32 := GenS1{F0: ("" + ("hello" + "hello"))}
    for fi33 := 0; fi33 < 8; fi33++ {
        if true {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int = (fi33 - fi33)
                fmt.Println(fi34, fi33)
                v36 := GenS2{F0: 32767, F1: 143}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        if (v3 > (v3 ^ v3)) {
            fmt.Println(v6, v31)
            for fi37 := 0; fi37 < 7; fi37++ {
                v38 := GenS2{F0: -32768, F1: 150}
                _ = fi37
                _ = v38
            }
        }
        switch (fi33 * fi33) {
        case 486:
            v6 = (0 == 1000)
        case -1:
            v1 = append(v1, 864)
        }
        _ = fi33
    }
    switch 32768 {
    case -9223372036854775808:
        v3 = 18446744073709551615
    default:
        if v31 {
            switch 134 {
            case 128:
                v3 -= v3
            case -87:
                v39 := GenS3{F0: 0}
                _ = v39
            }
        } else {
            v40 := v5[-921]
            _ = v40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
