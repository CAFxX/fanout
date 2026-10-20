package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 1
    var v3 int8 = -128
    v4 := make(map[int]string)
    var v5 int = 1
    var v6 int16 = 32767
    v7 := GenS1{F0: -57, F1: false, F2: 128}
    if (42 != 65535) {
        for fi8 := 0; fi8 < 10; fi8++ {
            v5 += (v5 % ((952) | 1))
            var v9 bool = (false || (true || true))
            var v10 int16 = (v6 + (-(v6)))
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: int16(v2), F1: (1 <= 18446744073709551615), F2: (v3 | (v3 + v3))}
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v6 = v6
            _ = fi13
        }
        v3 = v3
        v6 -= 145
        _ = fi12
    }
    fmt.Println(v3, v6)
    switch -1 {
    case 2147483647:
        fmt.Println(v2)
    case -291:
        if false {
            switch ("hello" + "foo bar") {
            case "":
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 uint8 = 255
                    v16 := v4[fi14]
                    v17 := GenS1{F0: (v3 ^ (v3 / ((v3) | 1))), F1: (((v16 + (v16 + v16)) + ("a" + "a")) <= v16), F2: v15}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
            v5 = v5
        }
    case 128:
        v6 = int16(v3)
    }
    v2 -= (v2 + v2)
    v18 := GenS2{F0: v6, F1: false, F2: (^(v3))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
