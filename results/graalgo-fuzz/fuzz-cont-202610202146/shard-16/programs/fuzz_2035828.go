package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 42
    var v2 int16 = 48
    v3 := []int{-2147483648, -329}
    var v4 float32 = 1.5
    var v5 uint16 = 2
    v6 := make(map[string]string)
    var v7 int8 = -120
    if ((1 != 1) && (2 >= 518)) {
        fmt.Println(v5, v1)
    }
    var v8 uint64 = (v1 % (((v1 - v1)) | 1))
    fmt.Println(v4, v1)
    v9 := GenS1{F0: true}
    v2 += v2
    var v10 int8 = 0
    v5 -= (v5 - 32767)
    var v11 uint8 = 89
    for fi12 := 0; fi12 < 1; fi12++ {
        delete(v6, ("hello" + ("\n\t" + (("\n\t" + "") + "hello"))))
        v13 := GenS1{F0: true}
        v5 += 952
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v2 = (v2 - v2)
        v15 := GenS1{F0: ((0 == 2) || false)}
        for fi16 := 0; fi16 < 17; fi16++ {
            v3 = append(v3, fi14)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v17 int32 = 42
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v2)
    var v18 int16 = v2
    for fi19 := 0; fi19 < 15; fi19++ {
        v1 = 4294967295
        _ = len(v3)
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v1, v18, fi19)
            _ = fi20
        }
        _ = fi19
    }
    var v21 int32 = 0
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            delete(v6, "\n\t")
            v10 -= (v7 << ((uint(v11)) & 63))
            _ = fi23
        }
        v11 = 2
        if (!((1 >= (v21 * v21)) && true)) {
            v6[("hello" + "")] = ("a" + ((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + ("hello" + "\n\t"))))
        } else {
            if (false && (v10 > 42)) {
                v24 := GenS1{F0: true}
                var v25 uint64 = (v8 - v1)
                v25 = (^(v1))
                _ = v24
                _ = v25
            } else {
                var v26 bool = (v1 != v8)
                _ = v26
            }
            if false {
                v21 -= v21
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint16 = v5
                    v3 = append(v3, (^(fi27)))
                    v29 := GenS1{F0: (!(977 == 564))}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := GenS2{F0: 9223372036854775807}
                var v32 int32 = int32(fi22)
                var v33 uint16 = (^(v5))
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        _ = fi22
    }
    v34 := v6[("\n\t" + "hello")]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
