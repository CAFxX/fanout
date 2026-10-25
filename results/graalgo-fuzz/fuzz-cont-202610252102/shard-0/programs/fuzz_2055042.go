package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 int8) string {
    v7 := GenS1{F0: 127, F1: 2147483648}
    var v8 int8 = p0
    v8 += -1
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := []int{32767}
    v2 := []string{"\n\t"}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 uint16 = 2
    var v6 int64 = 42
    v9 := GenS2{F0: 42, F1: 23}
    var v10 int64 = (-(v6))
    fmt.Println(v6)
    var v11 uint32 = 451
    var v12 bool = true
    v6 += -9223372036854775808
    v5 = v5
    v13 := GenS3{F0: (v11 << ((198) & 63)), F1: (127 != 0), F2: ("" + "foo bar")}
    fmt.Println(v6)
    v14 := GenS2{F0: 473, F1: 2}
    v15 := GenS4{F0: 68}
    if v12 {
        v11 = v11
    }
    fmt.Println(v6, v12, v5)
    if false {
        for fi16 := 0; fi16 < 7; fi16++ {
            v10 -= (v6 - -738)
            _ = fi16
        }
        if v12 {
            v11 += (v11 << ((825) & 63))
            for fi17 := 0; fi17 < 1; fi17++ {
                v4[fi17] = true
                for fi18 := 0; fi18 < 10; fi18++ {
                    _ = len(v2)
                    var v19 int = fi18
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 5; fi20++ {
                    _ = len(v4)
                    v6 -= (v10 - (v10 & v10))
                    _ = fi20
                }
                _ = fi17
            }
        } else {
            if v12 {
                var v21 uint8 = uint8(v5)
                var v22 uint32 = (v11 - (v11 << ((65535) & 63)))
                _ = v21
                _ = v22
            } else {
                var v23 uint32 = 2147483647
                var v24 int8 = 42
                for fi25 := 0; fi25 < 19; fi25++ {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v6 = v10
                        fmt.Println(fi25, v6)
                        _ = fi26
                    }
                    v27 := GenS3{F0: v23, F1: (!(!false)), F2: ("" + (("foo bar" + ("foo bar" + "a")) + "\n\t"))}
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v6 += (v6 & (^(v6)))
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        _ = fi28
                    }
                    _ = fi25
                    _ = v27
                }
                _ = v23
                _ = v24
            }
            var v29 uint16 = v5
            for fi30 := 0; fi30 < 17; fi30++ {
                fmt.Println(v5)
                if ("" <= ("\n\t" + "")) {
                    fmt.Println(v5, v29, v10)
                    v31 := GenS5{F0: "\n\t", F1: "\n\t"}
                    fmt.Println(v10, v5)
                    _ = v31
                }
                fmt.Println(v29, v10)
                _ = fi30
            }
            _ = v29
        }
    } else {
        v32 := GenS4{F0: 42}
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
