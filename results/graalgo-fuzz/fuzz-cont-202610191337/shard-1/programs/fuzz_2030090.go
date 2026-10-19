package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 uint16 = (v1 << ((0) & 63))
    var v3 int64 = -385
    if (uint64(v2) >= 369) {
        if (!(0 >= 4294967295)) {
            var v4 uint64 = 0
            v5 := []int{4294967295, -206}
            _ = v4
            _ = v5
        } else {
            if true {
                v2 -= v1
                v6 := GenS1{F0: -1, F1: 1, F2: "a"}
                v7 := GenS1{F0: 9223372036854775807, F1: 178, F2: "a"}
                _ = v6
                _ = v7
            }
        }
        switch 127 {
        case 32768:
            var v8 bool = false
            _ = v8
        default:
            var v9 int = 32768
            _ = v9
        }
    }
    var v10 uint = 1
    for fi11 := 0; fi11 < 20; fi11++ {
        v10 -= 48
        _ = fi11
    }
    if (127 >= 127) {
        v3 -= v3
        var v12 int8 = 0
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS1{F0: fi13, F1: 0, F2: "foo bar"}
            v10 += (^(v10))
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    fmt.Println(v1, v10)
    v15 := []int{}
    var v16 string = ("" + "foo bar")
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(v1, v16, v2)
        _ = fi17
    }
    switch -1 {
    case 32767:
        switch (("hello" + (((v16 + v16) + (((v16 + (v16 + (("hello" + v16) + (v16 + v16)))) + "hello") + ("foo bar" + (v16 + (v16 + v16))))) + v16)) + v16) {
        case "hello":
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS1{F0: fi18, F1: 255, F2: v16}
                _ = fi18
                _ = v19
            }
        case "foo bar":
            v1 = 1
        }
    case -396:
        switch "a" {
        case "foo bar":
            if ((204 == 0) || (65535 == uint32(v1))) {
                var v20 uint = 778
                v21 := GenS2{F0: 2147483647, F1: 2, F2: 0}
                _ = v20
                _ = v21
            }
        case "\n\t":
            fmt.Println(v10, v16)
        case "":
            for fi22 := 0; fi22 < 3; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (true && (!(75 == 1))) {
                        if len(v15) > 0 {
                            v15[0] = fi23
                        }
                        v3 = -1
                        var v24 int = 2147483647
                        _ = v24
                    }
                    _ = fi23
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi22
            }
        default:
            fmt.Println(v1)
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
