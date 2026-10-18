package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }

func main() {
    var v1 float32 = ((-2.25 * -2.25) * ((((1.5 + ((1e10 - 0.1) + 1.5)) - 3.14159) + 0.1) * ((1.5 - (0.1 - (0.1 * 3.14159))) - 0.0)))
    var v2 string = ""
    var v3 int64 = -1
    v4 := make(map[string]bool)
    var v5 uint8 = 62
    _ = len(v4)
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(v5, v2)
        _ = fi6
    }
    var v7 float32 = 1.5
    if true {
        var v8 uint16 = 1
        if false {
            if true {
                v5 = v5
                for fi9 := 0; fi9 < 7; fi9++ {
                    if (0 <= 311) {
                        delete(v4, v2)
                        var v10 uint64 = 1
                        for fi11 := 0; fi11 < 19; fi11++ {
                            var v12 uint16 = 478
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    } else {
                        v5 += 1
                    }
                    fmt.Println(v1)
                    fmt.Println(v2, v7)
                    _ = fi9
                }
            }
            v13 := []int{}
            v8 -= v8
            _ = v13
        }
        if (42 < 351) {
            switch (((v2 + (v2 + v2)) + "") + ((v2 + (v2 + (v2 + v2))) + (v2 + v2))) {
            case "":
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := []int{}
                    _ = fi14
                    _ = v15
                }
            case "foo bar":
                var v16 int8 = 127
                _ = v16
            }
            var v17 float32 = (v7 / (v7 / 0.1))
            _ = v17
        }
        _ = v8
    } else {
        fmt.Println(v1)
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v5 = 193
        _ = fi18
    }
    var v19 int16 = 32767
    var v20 int8 = 0
    for fi21 := 0; fi21 < 18; fi21++ {
        switch fi21 {
        case 127:
            v2 = (v2 + "\n\t")
        }
        v5 = (v5 % (((v5 + (v5 >> ((275) & 63)))) | 1))
        _ = fi21
    }
    fmt.Println(v1)
    fmt.Println(v20, v5, v2)
    v22 := GenS1{F0: 65535, F1: 255, F2: (114 == uint64(v5))}
    var v23 int16 = v19
    var v24 int16 = (v23 / ((v19) | 1))
    if true {
        for fi25 := 0; fi25 < 12; fi25++ {
            if true {
                v26 := []int{fi25, (fi25 & fi25), (^(fi25)), fi25}
                var v27 uint16 = 255
                var v28 uint32 = 2147483648
                _ = v26
                _ = v27
                _ = v28
            }
            v1 = float32(-2.25)
            v29 := GenS1{F0: 362, F1: 18446744073709551615, F2: true}
            _ = fi25
            _ = v29
        }
    } else {
        fmt.Println(v19, v5)
    }
    v4["a"] = (v19 > (-(v24)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
