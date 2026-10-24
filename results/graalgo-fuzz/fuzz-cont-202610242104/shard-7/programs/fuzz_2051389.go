package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 int8) bool {
    fmt.Println(p0)
    var v5 float64 = 3.14159
    v5 = v5
    var v6 uint16 = 2
    p0 = (p0 << ((4294967295) & 63))
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 int32 = -94
    var v4 uint8 = 41
    v3 = v3
    v3 = (v3 ^ (v3 * v3))
    var v7 int = 127
    var v8 int = 0
    var v9 uint16 = uint16(v4)
    switch 32767 {
    case -9223372036854775808:
        v9 = 32767
    }
    v3 -= -1
    var v10 int32 = (v3 * v3)
    switch ((("a" + ("hello" + "foo bar")) + "a") + (((("foo bar" + "") + "a") + "a") + "a")) {
    case "a":
        if false {
            v4 += v4
        }
    case "\n\t":
        v11 := GenS1{F0: 2}
        _ = v11
    case "hello":
        if (!true) {
            switch ("\n\t" + "a") {
            case "\n\t":
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 uint64 = 0
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v15 := []int{2147483647, (fi14 | 32768), (^(fi12)), v7}
                        v16 := GenS1{F0: 2}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if false {
                        v17 := []int{fi12, 32767, (v7 >> ((2) & 63)), 1}
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                }
            case "hello":
                if (!true) {
                    if false {
                        var v18 int32 = (v3 >> ((255) & 63))
                        v19 := []int{}
                        if len(v19) > 0 {
                            v19[0] = (^(v8))
                        }
                        _ = v18
                        _ = v19
                    } else {
                        if ((-(v7)) == (-(v8))) {
                            if (!(v8 >= v7)) {
                                var v20 uint64 = 172
                                v21 := GenS2{F0: 42, F1: 42, F2: v20}
                                _ = v20
                                _ = v21
                            }
                            var v22 uint32 = 638
                            for fi23 := 0; fi23 < 9; fi23++ {
                                switch int(v3) {
                                case 520:
                                    v24 := []int{v7, v8}
                                    _ = v24
                                }
                                v2[v22] = "hello"
                                _ = fi23
                            }
                            _ = v22
                        }
                        fmt.Println(v4, v7, v3)
                        v7 = 32767
                    }
                    v25 := GenS2{F0: 42, F1: 582, F2: uint64(v9)}
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v10 += v3
                if true {
                    switch "\n\t" {
                    case "a":
                        v8 -= (v8 + v7)
                    case "\n\t":
                        fmt.Println(v3, v7)
                    }
                }
                _ = fi26
            }
            var v27 int16 = 842
            _ = v27
        } else {
            var v28 float32 = 0.1
            _ = v28
        }
    default:
        v29 := GenS1{F0: 42}
        _ = v29
    }
    fmt.Println(v7, v4)
    v30 := GenS1{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
