package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 int32) int64 {
    var v6 bool = true
    switch 1 {
    case 9223372036854775807:
        v7 := GenS1{F0: 127, F1: "a", F2: (1 <= 1000)}
        _ = v7
    }
    _ = p0
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 83
    var v3 uint64 = 711
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint16 = 32767
        switch "" {
        case "\n\t":
            var v10 float32 = 0.1
            _ = v10
        }
        var v11 int32 = 1
        _ = fi8
        _ = v11
        _ = v9
    }
    var v12 uint8 = v2
    var v13 uint64 = 726
    var v14 float32 = 0.0
    v15 := GenS1{F0: (v12 - (v12 - (v2 & (v12 << ((51) & 63))))), F1: "hello", F2: (false || (194 < 1))}
    var v16 int8 = 127
    if false {
        v12 = v2
        if true {
            var v17 float64 = 1e10
            v17 = v17
            for fi18 := 0; fi18 < 19; fi18++ {
                if ((genFn0(-2147483648) == 42) || false) {
                    v19 := GenS1{F0: (v2 % ((v12) | 1)), F1: (((("hello" + "") + ("a" + "\n\t")) + "hello") + "\n\t"), F2: (v16 <= (v16 / ((v16) | 1)))}
                    v13 = (v13 + v3)
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v17 = (((0.0 + v17) + 1.5) / v17)
                        _ = fi20
                    }
                    _ = v19
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 float32 = v14
                    v23 := GenS1{F0: v12, F1: "foo bar", F2: true}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi18
            }
            _ = v17
        }
    }
    var v24 int8 = (v16 << (((v1 - (v1 * v1))) & 63))
    delete(v4, 2147483647)
    v14 -= v14
    v14 -= ((3.14159 / v14) + 1e10)
    if ((!(-2.25 != float64((v14 + (v14 - (v14 + (0.1 / v14))))))) && true) {
        for fi25 := 0; fi25 < 2; fi25++ {
            if false {
                var v26 uint32 = 65535
                var v27 int32 = 662
                _ = v26
                _ = v27
            } else {
                if (v14 < (v14 - v14)) {
                    v16 -= v24
                }
                var v28 uint16 = 42
                _ = v28
            }
            v2 -= 0
            _ = fi25
        }
    } else {
        var v29 bool = false
        var v30 uint16 = 590
        v30 = 556
        _ = v29
        _ = v30
    }
    var v31 bool = (2147483647 != -2147483648)
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := []int{fi32}
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: 174, F1: "hello", F2: (1 > 1000)}
    v35 := GenS1{F0: 202, F1: "a", F2: (!((("a" + ("a" + "\n\t")) + "foo bar") == "a"))}
    v13 = 994
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
