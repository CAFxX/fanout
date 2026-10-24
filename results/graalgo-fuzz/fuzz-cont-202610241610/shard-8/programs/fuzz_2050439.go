package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint64, p2 int16) uint32 {
    p0 = (true && false)
    v7 := GenS1{F0: -2147483648, F1: 32767, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"foo bar", ((("a" + "") + (("\n\t" + (("foo bar" + ("" + ("a" + ("" + "a")))) + ("a" + "a"))) + "\n\t")) + ("foo bar" + "hello"))}
    v3 := []int{}
    v4 := []string{(("\n\t" + "hello") + "hello"), ("" + "foo bar")}
    var v5 uint8 = 1
    v6 := []uint32{4294967295, 2147483647}
    switch -2147483648 {
    case 32767:
        var v8 int = 259
        _ = v8
    case -2147483648:
        v5 -= (v5 << ((4294967295) & 63))
    case -1:
        var v9 string = (("\n\t" + "\n\t") + "foo bar")
        _ = v9
    default:
        v10 := v1[940]
        _ = v10
    }
    if (-1 >= 42) {
        if (-513 < -9223372036854775808) {
            v11 := GenS2{F0: "hello"}
            if true {
                v5 -= 255
                delete(v1, 2147483648)
            } else {
                v5 -= 42
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int32 = int32(v5)
                    for fi14 := 0; fi14 < 10; fi14++ {
                        delete(v1, 0)
                        switch (fi12 ^ -9223372036854775808) {
                        case 501:
                            _ = len(v4)
                        case -1:
                            var v15 int16 = 42
                            _ = v15
                        }
                        _ = fi14
                    }
                    fmt.Println(v13, v5)
                    _ = fi12
                    _ = v13
                }
                var v16 bool = ("\n\t" != (("foo bar" + "hello") + "hello"))
                _ = v16
            }
            _ = v11
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(fi17, fi17)
            for fi18 := 0; fi18 < 12; fi18++ {
                if true {
                    v5 = v5
                    v4 = append(v4, "a")
                } else {
                    var v19 uint32 = 457
                    v20 := GenS1{F0: (^(fi18)), F1: 65535, F2: fi18}
                    v19 += v19
                    var v21 int16 = -940
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v5 += (v5 + 255)
                _ = fi18
            }
            _ = fi17
        }
    } else {
        fmt.Println(v5)
        for fi22 := 0; fi22 < 11; fi22++ {
            v5 += v5
            v23 := GenS1{F0: fi22, F1: 65535, F2: fi22}
            _ = len(v3)
            _ = fi22
            _ = v23
        }
        fmt.Println(v5)
    }
    v4 = append(v4, ("a" + "a"))
    var v24 int32 = 0
    fmt.Println(v5, v24)
    fmt.Println(v5, v24)
    v5 += v5
    v25 := GenS1{F0: 398, F1: 642, F2: -9223372036854775808}
    v26 := GenS2{F0: "hello"}
    v27 := GenS1{F0: int(v24), F1: 255, F2: 2147483647}
    switch 1 {
    case 806:
        if (true || (true || true)) {
            _ = len(v3)
        } else {
            v1[2] = "\n\t"
        }
    }
    _ = len(v1)
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint = 255
        if (1 != 127) {
            fmt.Println(v5, v24)
        } else {
            var v30 uint16 = 458
            var v31 int32 = (v24 + v24)
            _ = v30
            _ = v31
        }
        v5 = v5
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
