package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 string; F1 int; F2 string }
func genFn0(p0 int8, p1 uint64, p2 string) int8 {
    if (true || true) {
        v6 := GenS1{F0: -67, F1: 1, F2: -1}
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint8 = 153
            _ = fi7
            _ = v8
        }
        var v9 int32 = 321
        _ = v6
        _ = v9
    }
    p1 = p1
    v10 := GenS1{F0: -1, F1: 0, F2: -580}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 int8 = -62
    v2 := make(map[string]int)
    var v3 int32 = 0
    var v4 float32 = (1e10 / 1.0)
    v5 := make(map[string]uint32)
    v11 := GenS2{F0: 65535, F1: -1, F2: v1}
    var v12 int8 = v1
    v13 := v5["hello"]
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: 42, F1: 18446744073709551615, F2: -713}
        for fi16 := 0; fi16 < 17; fi16++ {
            if (v4 >= v4) {
                var v17 bool = false
                _ = len(v2)
                if v17 {
                    if v17 {
                        v13 = v13
                        var v18 uint64 = 0
                        _ = v18
                    }
                    var v19 int8 = (v12 + int8(v3))
                    _ = v19
                }
                v1 = genFn0(100, 2, "foo bar")
                _ = v17
            } else {
                switch "foo bar" {
                case "foo bar":
                    var v20 int32 = (v3 << ((42) & 63))
                    _ = v20
                case "":
                    var v21 int64 = 0
                    _ = v21
                }
            }
            v22 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: 9223372036854775807}
            v23 := GenS3{F0: ("foo bar" + ((("hello" + "hello") + "") + (("" + ((("hello" + ("\n\t" + "foo bar")) + "hello") + ("hello" + ("" + "a")))) + ("a" + "\n\t")))), F1: 2147483647, F2: ("hello" + ("\n\t" + "a"))}
            _ = fi16
            _ = v22
            _ = v23
        }
        var v24 int16 = 32767
        _ = fi14
        _ = v15
        _ = v24
    }
    switch "\n\t" {
    case "":
        v12 += v1
    case "hello":
        v1 = int8(v3)
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 float32 = float32(1e10)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    v12 = v12
    var v28 uint8 = 131
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
