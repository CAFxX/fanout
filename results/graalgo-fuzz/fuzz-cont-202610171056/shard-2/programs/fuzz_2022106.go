package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 uint16 = 2
    v3 -= 32767
    v3 = (v3 * v3)
    v3 -= 32767
    var v4 uint16 = v3
    if (!false) {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: 1}
            _ = fi5
            _ = v6
        }
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 0}
                _ = fi7
                _ = v8
            }
            if (0.0 >= 0.1) {
                fmt.Println(v3, v3)
                fmt.Println(v4, v4)
            }
            fmt.Println(v4, v4)
        } else {
            if (32767 < v4) {
                var v9 int = 127
                v10 := GenS1{F0: 245}
                _ = v10
                _ = v9
            }
            fmt.Println(v4, v3)
            v4 -= (v3 / ((v4) | 1))
        }
    } else {
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := v2[""]
                _ = fi11
                _ = v12
            }
            v13 := []int{}
            var v14 int = 9223372036854775807
            _ = v13
            _ = v14
        }
        if false {
            for fi15 := 0; fi15 < 8; fi15++ {
                v4 = v3
                if (!(true || (4294967295 < 0))) {
                    fmt.Println(v3)
                } else {
                    v3 -= v3
                    var v16 int32 = -452
                    _ = v16
                }
                v4 += v4
                _ = fi15
            }
            var v17 bool = true
            _ = v17
        } else {
            var v18 uint = 42
            _ = v18
        }
    }
    if ((v3 % ((v4) | 1)) <= (v4 + v3)) {
        var v19 int = 9223372036854775807
        _ = v19
    }
    v20 := GenS1{F0: 128}
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int8 = 100
        if ((0 == 850) && true) {
            switch fi21 {
            case 4294967295:
                v23 := GenS2{F0: (v3 / ((1) | 1)), F1: 42}
                _ = v23
            case -391:
                var v24 string = ("" + ("foo bar" + ""))
                _ = v24
            case 2147483647:
                fmt.Println(v22, v3)
            default:
                var v25 uint8 = 128
                _ = v25
            }
        }
        fmt.Println(v4, v22, fi21)
        _ = fi21
        _ = v22
    }
    var v26 uint = 752
    var v27 uint32 = 2
    delete(v2, ((("foo bar" + ("foo bar" + "foo bar")) + "") + "\n\t"))
    if false {
        fmt.Println(v3)
        fmt.Println(v26, v26)
    }
    v26 -= (v26 & 1)
    v28 := GenS2{F0: (^(v3)), F1: v27}
    fmt.Println(v26, v4, v4)
    if false {
        switch 2147483647 {
        case 127:
            switch "foo bar" {
            case "":
                switch (("" + "foo bar") + "") {
                case "hello":
                    var v29 bool = (1.5 <= (-2.25 + (1.5 / 1e10)))
                    _ = v29
                default:
                    var v30 uint8 = 0
                    _ = v30
                }
            case "hello":
                delete(v1, 340)
            case "a":
                v26 += v26
            default:
                if (uint16(v26) > v3) {
                    var v31 uint = v26
                    v4 += 17
                    _ = v31
                }
            }
        default:
            v27 = (v27 << (((v26 / ((368) | 1))) & 63))
        }
        v32 := GenS1{F0: 0}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
