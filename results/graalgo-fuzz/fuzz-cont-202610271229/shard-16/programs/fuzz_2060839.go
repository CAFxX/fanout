package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi5
        _ = v6
    }
    if (127 > 128) {
        v7 := GenS1{F0: 255, F1: (-749 < -164), F2: 0}
        var v8 uint = 18446744073709551615
        v8 -= 255
        _ = v7
        _ = v8
    }
    var v9 int32 = -815
    _ = v9
    return -32768
}
func genFn1(p0 string, p1 uint8) int8 {
    var v10 int32 = 2147483647
    var v11 int16 = -32768
    if false {
        var v12 uint8 = (p1 << ((241) & 63))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 uint16 = 65535
    var v4 float32 = 3.14159
    switch -2147483648 {
    case 158:
        if (!(!(true && (true && false)))) {
            var v13 bool = ((-9223372036854775808 <= -249) || true)
            v14 := GenS1{F0: 18446744073709551615, F1: (196 == 2), F2: 918}
            v1[32768] = "a"
            _ = v13
            _ = v14
        }
    case -818:
        for fi15 := 0; fi15 < 10; fi15++ {
            v1[(-(fi15))] = v2
            _ = fi15
        }
    }
    v16 := GenS1{F0: 18446744073709551615, F1: true, F2: 18446744073709551615}
    fmt.Println(v4, v3)
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS2{F0: 0, F1: 127, F2: 18446744073709551615}
                var v20 int = (-(fi17))
                v4 -= v4
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            v4 = float32(0.1)
        }
        _ = fi17
    }
    v4 = float32(((1e10 - float64(((3.14159 * float32(1.5)) - v4))) - (1.5 * float64(-2.25))))
    fmt.Println(v4, v4, v2)
    if (65535 <= 65535) {
        if false {
            v21 := GenS3{F0: v3}
            _ = v21
        }
        v3 += v3
        var v22 float64 = 0.1
        _ = v22
    }
    v3 -= (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
