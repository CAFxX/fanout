package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64, p1 bool, p2 string) int {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 -= (p0 % ((90) | 1))
        v6 := GenS1{F0: 18446744073709551615}
        if p1 {
            var v7 int32 = 0
            _ = v7
        } else {
            p2 = (("foo bar" + p2) + "a")
            for fi8 := 0; fi8 < 9; fi8++ {
                p2 = ((p2 + (p2 + (p2 + (p2 + ((p2 + ("\n\t" + p2)) + ("\n\t" + p2)))))) + (("\n\t" + p2) + "a"))
                v9 := []int{-526, (fi5 * (fi5 % (((fi8 * (fi5 - 32767))) | 1))), fi8, fi5}
                var v10 uint32 = 2147483648
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi5
        _ = v6
    }
    var v11 int = -821
    for fi12 := 0; fi12 < 8; fi12++ {
        p1 = p1
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return v11
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 0
    var v3 uint8 = 255
    v4 := make(map[int]bool)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 42}
        var v15 uint8 = (v3 | v3)
        _ = len(v4)
        _ = fi13
        _ = v14
        _ = v15
    }
    v1 = 2147483647
    var v16 bool = true
    _ = len(v4)
    v17 := GenS1{F0: 676}
    fmt.Println(v1, v3, v16)
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := []int{fi18, fi18, (fi18 << ((4294967295) & 63))}
        switch (fi18 & (fi18 | 4294967295)) {
        case 128:
            v16 = (v1 != 42)
        case -639:
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int8 = 127
                var v22 uint8 = 15
                fmt.Println(v3, v1)
                _ = fi20
                _ = v21
                _ = v22
            }
        default:
            v4[fi18] = (v2 == (v2 & 2))
        }
        v23 := GenS1{F0: uint64(v3)}
        _ = fi18
        _ = v19
        _ = v23
    }
    switch ("a" + "hello") {
    case "hello":
        if (-1 > 1) {
            var v24 uint = 255
            switch "\n\t" {
            case "":
                for fi25 := 0; fi25 < 4; fi25++ {
                    if true {
                        var v26 uint32 = 484
                        _ = v26
                    }
                    _ = fi25
                }
            default:
                if v16 {
                    v27 := GenS1{F0: 18446744073709551615}
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 int64 = -9223372036854775808
                        for fi30 := 0; fi30 < 5; fi30++ {
                            if ((0.1 + 1e10) >= (0.1 - (-2.25 * (3.14159 + -2.25)))) {
                                var v31 int64 = (-(v29))
                                v32 := []int{-870, (-(fi30)), fi28, fi28}
                                _ = v31
                                _ = v32
                            } else {
                                v33 := []int{}
                                fmt.Println(v1)
                                var v34 uint8 = uint8(v2)
                                _ = v33
                                _ = v34
                            }
                            _ = fi30
                        }
                        _ = fi28
                        _ = v29
                    }
                    for fi35 := 0; fi35 < 6; fi35++ {
                        var v36 uint32 = 29
                        fmt.Println(v16, v16)
                        _ = fi35
                        _ = v36
                    }
                    _ = v27
                }
            }
            v37 := GenS1{F0: 18446744073709551615}
            _ = v24
            _ = v37
        }
    case "":
        if (65535 >= 4294967295) {
            v3 = v3
            var v38 int8 = -42
            _ = v38
        } else {
            var v39 int64 = -1
            v3 += 1
            _ = v39
        }
    }
    v40 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
