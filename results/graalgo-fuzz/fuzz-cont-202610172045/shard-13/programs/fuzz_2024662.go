package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int, p1 int32, p2 int) int16 {
    for fi5 := 0; fi5 < 15; fi5++ {
        switch (p2 * (p2 ^ fi5)) {
        case 32767:
            p1 = 42
        }
        _ = fi5
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 576
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -535
    var v3 int32 = (v2 & -1)
    v4 := make(map[string]string)
    fmt.Println(v3, v2)
    v6 := []int{32768}
    if (0 > 2) {
        switch 146 {
        case -2147483648:
            var v7 float32 = float32(-2.25)
            _ = v7
        case 2147483647:
            if (!(true && true)) {
                for fi8 := 0; fi8 < 13; fi8++ {
                    v3 = (v3 & -2147483648)
                    var v9 uint8 = 1
                    _ = fi8
                    _ = v9
                }
                switch "foo bar" {
                case "":
                    var v10 string = "\n\t"
                    _ = v10
                case "a":
                    if (!true) {
                        var v11 int64 = 309
                        _ = v11
                    } else {
                        if ((!false) || (((float32(0.1) / 0.1) / float32(1e10)) <= 0.1)) {
                            var v12 string = ("foo bar" + "\n\t")
                            v13 := GenS1{F0: 680}
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint64 = 1
                            _ = len(v1)
                            if (457 < 255) {
                                var v15 bool = (!(false && false))
                                _ = v15
                            }
                            _ = v14
                        }
                        v2 = v3
                        v3 = (v3 & v2)
                    }
                }
                fmt.Println(v2)
            }
        }
    }
    v2 = 754
    v16 := GenS1{F0: 18446744073709551615}
    v2 -= 2147483647
    for fi17 := 0; fi17 < 14; fi17++ {
        if len(v6) > 0 {
            v6[0] = (fi17 / ((127) | 1))
        }
        fmt.Println(v2, fi17)
        for fi18 := 0; fi18 < 8; fi18++ {
            v6 = append(v6, fi17)
            _ = fi18
        }
        _ = fi17
    }
    if false {
        var v19 int16 = 0
        _ = len(v4)
        switch 0 {
        case 672:
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := v1[2]
                var v22 float32 = float32(-2.25)
                fmt.Println(fi20)
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = v19
    }
    v23 := GenS1{F0: 18446744073709551615}
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint16 = 65535
        var v26 uint16 = v25
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if true {
                var v27 uint32 = 21
                _ = v27
            }
            v25 = 40
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = len(v6)
    v28 := GenS1{F0: 4294967295}
    fmt.Println(v3, v3)
    v29 := GenS1{F0: 964}
    var v30 uint64 = 892
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v6
}
