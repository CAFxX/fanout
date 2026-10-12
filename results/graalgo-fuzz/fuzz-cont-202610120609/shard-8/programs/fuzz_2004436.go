package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 bool = (604 != 4294967295)
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: fi4, F1: -32768}
        v6 := []int{42, 32767, fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: int16(v2)}
    v2 += v2
    v2 = v2
    if v3 {
        var v8 uint16 = 20
        for fi9 := 0; fi9 < 13; fi9++ {
            delete(v1, 2)
            var v10 uint64 = 4294967295
            v11 := GenS1{F0: 9223372036854775807, F1: -107}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        var v12 uint16 = 884
        _ = v12
    }
    var v13 string = ("\n\t" + "foo bar")
    v3 = v3
    for fi14 := 0; fi14 < 8; fi14++ {
        switch fi14 {
        case -2147483648:
            var v15 string = v13
            _ = v15
        }
        v2 += v2
        _ = fi14
    }
    var v16 int8 = -35
    v16 = v16
    v2 = (v2 << ((255) & 63))
    switch (((("foo bar" + "\n\t") + "\n\t") + ("" + v13)) + (("hello" + (v13 + (((v13 + (v13 + v13)) + ((v13 + "a") + v13)) + v13))) + "hello")) {
    case "a":
        var v17 bool = (0 == 33)
        _ = v17
    case "":
        if (!((!(v3 && (65535 >= 1))) || (!(false || ((-2147483648 < 42) && v3))))) {
            var v18 float64 = 0.0
            for fi19 := 0; fi19 < 20; fi19++ {
                _ = len(v1)
                _ = fi19
            }
            if (true || v3) {
                fmt.Println(v13, v2, v16)
                fmt.Println(v13, v18, v3)
            } else {
                v18 += ((v18 - float64(1.5)) + v18)
                if ((v16 > -1) && (536 <= 2)) {
                    v20 := GenS2{F0: -1}
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 int32 = 1
                        _ = fi21
                        _ = v22
                    }
                    v13 = "\n\t"
                    _ = v20
                } else {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS1{F0: 32767, F1: -456}
                        _ = fi23
                        _ = v24
                    }
                }
                delete(v1, 2147483648)
            }
            _ = v18
        }
    }
    if v3 {
        var v25 uint32 = 212
        var v26 int8 = (-(v16))
        _ = v25
        _ = v26
    }
    var v27 uint64 = 990
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v7
}
