package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
func genFn0() int16 {
    if (4294967295 != 127) {
        for fi5 := 0; fi5 < 6; fi5++ {
            if true {
                var v6 int8 = 42
                var v7 uint = 953
                v8 := []int{(fi5 / (((fi5 * -831)) | 1)), fi5, fi5}
                _ = v6
                _ = v7
                _ = v8
            } else {
                var v9 uint64 = 0
                if false {
                    if (!(0 < 128)) {
                        v10 := GenS1{F0: 4294967295, F1: 0, F2: (v9 << ((524) & 63))}
                        v9 = 1
                        _ = v10
                    }
                    var v11 uint8 = 2
                    _ = v11
                }
                _ = v9
            }
            _ = fi5
        }
        var v12 uint64 = 478
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{}
        switch 1 {
        case 1:
            v15 := GenS2{F0: 1, F1: fi13, F2: 65535}
            _ = v15
        case -222:
            switch ((("\n\t" + "foo bar") + "foo bar") + ("\n\t" + "\n\t")) {
            case "foo bar":
                var v16 float32 = (1e10 - (((float32(((1e10 * (1e10 + 1e10)) - (3.14159 + 0.1))) + (1.5 - 1e10)) - 0.1) * ((3.14159 + (-2.25 - 3.14159)) - (0.0 - (3.14159 / ((1.5 / 0.1) + (((3.14159 / -2.25) - (1e10 - float32((0.1 + (-2.25 * 1e10))))) + 3.14159)))))))
                _ = v16
            }
        }
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    var v17 string = (("a" + ("" + "\n\t")) + "")
    _ = v17
    return 0
}
func genFn1(p0 uint8) bool {
    fmt.Println(p0)
    v18 := []int{}
    _ = len(v18)
    _ = p0
    _ = v18
    return false
}

func main() {
    var v1 int16 = 42
    var v2 uint64 = 441
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v1 -= 32767
    var v19 int8 = -128
    v20 := []int{}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println(v1, v1, v19)
    for fi21 := 0; fi21 < 17; fi21++ {
        v1 = v1
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint16 = 32767
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if false {
        if ((0 > 2147483647) && true) {
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v19)
                _ = fi24
            }
            v25 := v3["hello"]
            _ = v25
        } else {
            v19 += v19
            v1 = (v1 / ((genFn0()) | 1))
        }
        v26 := GenS2{F0: (^(v19)), F1: -9223372036854775808, F2: 0}
        _ = v26
    }
    var v27 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
