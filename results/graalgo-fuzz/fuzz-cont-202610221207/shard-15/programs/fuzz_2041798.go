package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int, p1 int16) uint32 {
    v5 := []int{(p0 + p0)}
    v6 := GenS1{F0: "foo bar"}
    v7 := GenS1{F0: ("" + "\n\t")}
    fmt.Println(p1)
    if true {
        switch ((("foo bar" + "foo bar") + (("foo bar" + "a") + ((("hello" + "") + "hello") + "a"))) + "hello") {
        case "foo bar":
            for fi8 := 0; fi8 < 14; fi8++ {
                if (154 > -2147483648) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi8
            }
        case "":
            if true {
                switch ("" + "") {
                case "hello":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        _ = len(v5)
                        p1 -= 1000
                        p1 = 1000
                        _ = fi9
                    }
                case "a":
                    var v10 float64 = 0.0
                    _ = v10
                default:
                    switch p0 {
                    case -149:
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            for fi11 := 0; fi11 < 4; fi11++ {
                                var v12 int32 = -1
                                p0 += (p0 & p0)
                                _ = fi11
                                _ = v12
                            }
                            fmt.Println(p1, p0)
                        }
                    }
                }
                switch 245 {
                case 2147483647:
                    var v13 int64 = 42
                    _ = v13
                }
            }
        default:
            var v14 uint32 = 42
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint32 = 42
    var v3 bool = false
    var v4 uint64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += (v2 - (v2 << ((1) & 63)))
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            if v3 {
                var v17 uint8 = 1
                _ = v17
            }
            v18 := GenS1{F0: "\n\t"}
            v19 := GenS1{F0: ("foo bar" + "a")}
            _ = fi16
            _ = v18
            _ = v19
        }
        var v20 string = ""
        _ = fi15
        _ = v20
    }
    var v21 uint8 = 181
    fmt.Println(v3, v21)
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(v3, fi22)
        if true {
            fmt.Println(v4, v4, v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v2 += v2
            if ((v4 >> ((0) & 63)) <= 215) {
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := GenS1{F0: "\n\t"}
                    v21 = v21
                    _ = fi23
                    _ = v24
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v25 int32 = 1
            _ = v25
        }
        _ = fi22
    }
    if true {
        if len(v1) > 0 {
            v1[0] = genFn0(-1, 1000)
        }
        switch ("foo bar" + "") {
        case "hello":
            var v26 bool = false
            _ = v26
        case "a":
            for fi27 := 0; fi27 < 4; fi27++ {
                v4 -= (v4 & (v4 * v4))
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 uint16 = 0
                    var v30 int32 = 77
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v4 = (^(v4))
                _ = fi27
            }
        case "foo bar":
            var v31 uint8 = 127
            _ = v31
        default:
            var v32 uint32 = v2
            _ = v32
        }
    } else {
        v3 = (0.1 <= (0.0 - 3.14159))
    }
    fmt.Println(v2, v2, v21)
    if v3 {
        v1 = append(v1, v2)
        _ = len(v1)
    }
    v2 = (^(v2))
    v2 = (v2 | 0)
    v2 -= v2
    for fi33 := 0; fi33 < 3; fi33++ {
        if false {
            var v34 int16 = int16(v2)
            _ = v34
        } else {
            var v35 float64 = -2.25
            fmt.Println(fi33, v4, v2)
            _ = len(v1)
            _ = v35
        }
        v36 := GenS1{F0: (("a" + "\n\t") + ("a" + "a"))}
        _ = fi33
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
