package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 string) int8 {
    v4 := []int{-83, 128, 32768, 4294967295}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int8 = 12
    var v2 uint64 = 1
    var v3 string = ((("" + "a") + ("\n\t" + "\n\t")) + ("\n\t" + "hello"))
    v5 := GenS1{F0: 434, F1: "foo bar", F2: 255}
    var v6 string = v3
    switch -443 {
    case 2147483647:
        var v7 uint16 = 65535
        _ = v7
    case 42:
        var v8 uint8 = 2
        _ = v8
    case 9223372036854775807:
        var v9 bool = (!false)
        _ = v9
    }
    if (true || false) {
        v6 = "hello"
        if (true && true) {
            v3 = v6
            var v10 int8 = (v1 >> ((65535) & 63))
            _ = v10
        }
    }
    v1 -= (^(v1))
    var v11 uint16 = 0
    fmt.Println(v6)
    v2 -= 956
    v1 = (-(v1))
    if (640 < 0) {
        v2 = (v2 | (v2 % (((v2 + v2)) | 1)))
        v3 = v6
        if true {
            var v12 int16 = 154
            v13 := GenS1{F0: 9223372036854775807, F1: v3, F2: v11}
            _ = v12
            _ = v13
        }
    } else {
        var v14 float32 = float32(1.5)
        if false {
            v15 := []int{}
            if true {
                v16 := GenS1{F0: -1, F1: "foo bar", F2: (v11 / (((v11 | (v11 << ((255) & 63)))) | 1))}
                v17 := GenS1{F0: 0, F1: v6, F2: (v11 * (^(v11)))}
                fmt.Println(v14, v11, v6)
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v3 = v3
                    if len(v15) > 0 {
                        v15[0] = 880
                    }
                    v19 := GenS1{F0: int64(v11), F1: ("hello" + (((((("a" + v6) + v3) + ("foo bar" + (v6 + v6))) + "") + "") + "\n\t")), F2: (v11 - v11)}
                    _ = fi18
                    _ = v19
                }
                v15 = append(v15, -2147483648)
                if false {
                    fmt.Println(v14)
                }
            }
            v14 = 1e10
            _ = v15
        } else {
            v20 := GenS1{F0: -9223372036854775808, F1: "hello", F2: 32767}
            _ = v20
        }
        _ = v14
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
