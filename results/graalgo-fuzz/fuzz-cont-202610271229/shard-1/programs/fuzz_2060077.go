package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 string) uint64 {
    p0 = "hello"
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            var v7 int16 = -32768
            _ = v7
        } else {
            fmt.Println(fi6, p0)
            v8 := []int{fi6, (-(fi6))}
            var v9 uint16 = 0
            _ = v8
            _ = v9
        }
        if ((false && false) || true) {
            fmt.Println(p0)
            v10 := GenS1{F0: (0 <= 1)}
            if (127 <= (^(fi6))) {
                p0 = "\n\t"
            }
            _ = v10
        }
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int8 = 42
        if false {
            fmt.Println(v12, fi11, v12)
            for fi13 := 0; fi13 < 17; fi13++ {
                switch (p0 + "\n\t") {
                case "foo bar":
                    switch 0 {
                    case -9223372036854775808:
                        var v14 int8 = (v12 << ((0) & 63))
                        _ = v14
                    case 2147483647:
                        v12 = 42
                    case 128:
                        var v15 uint = 697
                        _ = v15
                    default:
                        if true {
                            var v16 int32 = -2147483648
                            var v17 int32 = (v16 >> ((642) & 63))
                            var v18 uint8 = 1
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            v19 := []int{(fi13 ^ 4294967295)}
                            v12 -= v12
                            _ = v19
                        }
                    }
                default:
                    switch "foo bar" {
                    case "a":
                        v20 := GenS1{F0: (!false)}
                        _ = v20
                    case "foo bar":
                        v21 := GenS2{F0: 1, F1: false, F2: 0}
                        _ = v21
                    default:
                        p0 = ""
                    }
                }
                v12 = v12
                var v22 uint64 = 42
                v12 = int8(fi13)
                _ = fi13
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int = fi23
                _ = fi23
                _ = v24
            }
            v12 -= v12
            for fi25 := 0; fi25 < 17; fi25++ {
                p0 = (p0 + "foo bar")
                v26 := []int{}
                v12 -= 1
                _ = fi25
                _ = v26
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []int{-2147483648}
    var v4 int16 = 32767
    var v5 int64 = 745
    switch "a" {
    case "\n\t":
        if true {
            v5 = 9223372036854775807
        }
    case "hello":
        delete(v1, 32767)
    case "a":
        switch 0 {
        case 1:
            _ = len(v2)
        case -20:
            v4 -= (v4 ^ (v4 << ((42) & 63)))
        }
    }
    v3 = append(v3, int(v4))
    var v27 uint = 1
    var v28 uint32 = 0
    if ((504 == 0) || true) {
        for fi29 := 0; fi29 < 12; fi29++ {
            v4 -= 1
            v27 = v27
            v27 = 212
            _ = fi29
        }
        v28 = v28
    }
    switch "a" {
    case "":
        switch 2147483647 {
        case 1:
            if true {
                _ = len(v2)
                var v30 uint64 = 464
                _ = v30
            } else {
                if len(v3) > 0 {
                    v3[0] = 774
                }
                v5 += (^(v5))
                v31 := GenS1{F0: (!(100 == 127))}
                _ = v31
            }
        case -649:
            for fi32 := 0; fi32 < 16; fi32++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 uint64 = 0
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
        }
    default:
        v35 := GenS1{F0: true}
        _ = v35
    }
    if (!false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v36 int32 = 0
    var v37 bool = false
    v38 := GenS3{F0: v5, F1: -1}
    if len(v3) > 0 {
        v3[0] = 827
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
