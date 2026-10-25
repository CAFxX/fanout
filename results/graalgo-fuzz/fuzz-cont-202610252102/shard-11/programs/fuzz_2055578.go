package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint32 = 0
        fmt.Println(fi5, v6)
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint = 706
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint16 = 42
        switch fi8 {
        case 128:
            var v10 bool = (!(2 >= 2))
            _ = v10
        }
        fmt.Println(v9, fi8)
        _ = fi8
        _ = v9
    }
    _ = v7
    return 32767
}
func genFn1(p0 int) bool {
    var v11 uint64 = 1
    v12 := GenS1{F0: (v11 / ((939) | 1))}
    v11 += v11
    v11 += (v11 - (v11 + v11))
    _ = p0
    _ = v11
    _ = v12
    return (0 <= 1)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -87
    v3 := make(map[uint32]bool)
    v4 := []int{-196, 128, 32768, -35}
    if true {
        _ = len(v1)
        _ = len(v4)
    }
    switch -1 {
    case 127:
        fmt.Println(v2)
    case 2147483647:
        _ = len(v3)
    case -877:
        if (!(-858 != 1)) {
            for fi13 := 0; fi13 < 17; fi13++ {
                switch -410 {
                case 9223372036854775807:
                    var v14 uint32 = 1
                    _ = v14
                case -467:
                    switch -783 {
                    case 128:
                        v2 += v2
                    }
                }
                var v15 float32 = -2.25
                _ = fi13
                _ = v15
            }
            v2 += (v2 << ((2) & 63))
            if (!false) {
                var v16 string = "a"
                _ = v16
            }
        } else {
            fmt.Println(v2)
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int8 = v2
                _ = fi17
                _ = v18
            }
            v2 -= (^(v2))
        }
    }
    var v19 bool = true
    v19 = v19
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS1{F0: 460}
            var v24 uint8 = 128
            v25 := GenS1{F0: 18446744073709551615}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        v4 = append(v4, (fi21 - (fi21 / (((fi21 / ((fi21) | 1))) | 1))))
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 float64 = 1.5
            _ = fi26
            _ = v27
        }
        _ = fi21
    }
    switch (("foo bar" + "foo bar") + ("\n\t" + "")) {
    case "a":
        fmt.Println(v20, v19, v19)
    }
    v20 -= (-(v20))
    _ = len(v3)
    var v28 int8 = v2
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
