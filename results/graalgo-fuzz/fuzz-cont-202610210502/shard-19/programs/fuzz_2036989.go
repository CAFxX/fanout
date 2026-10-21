package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint32; F2 int32 }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 65535
    var v3 int = 0
    v4 := []int{(v3 - v3), v3, (v3 * 803), 4294967295}
    if false {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: v1, F1: -1}
            v7 := GenS2{F0: 18446744073709551615, F1: "a"}
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 -= (v1 ^ v1)
            var v9 bool = true
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        v10 := GenS3{F0: 4294967295, F1: (v2 + v2), F2: (v2 >> ((16) & 63))}
        _ = v10
    } else {
        v3 -= 127
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            if false {
                var v13 uint = 0
                fmt.Println(v13, v13, v3)
                v14 := GenS2{F0: v13, F1: "foo bar"}
                v13 += v13
                _ = v13
                _ = v14
            } else {
                v15 := GenS4{F0: 1000}
                v3 += (v3 << ((2) & 63))
                _ = v15
            }
            _ = fi12
        }
        v3 = fi11
        _ = fi11
    }
    var v16 uint16 = v2
    v1 -= (v1 | v1)
    v1 -= -916
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint32 = 960
            v1 -= -754
            v3 = fi18
            _ = fi18
            _ = v19
        }
        fmt.Println(fi17, v3, v1)
        _ = fi17
    }
    if false {
        v16 -= v2
        if len(v4) > 0 {
            v4[0] = v3
        }
        v4 = append(v4, (v3 & v3))
    } else {
        switch (^(v3)) {
        case 2147483647:
            var v20 int32 = (v1 - 0)
            _ = v20
        case 32768:
            v16 -= v16
        case 0:
            for fi21 := 0; fi21 < 20; fi21++ {
                v1 += (v1 + v1)
                _ = fi21
            }
        }
        if (("foo bar" + "\n\t") < "hello") {
            if false {
                v2 = 606
                v22 := GenS5{F0: "\n\t"}
                _ = v22
            } else {
                fmt.Println(v16)
                if len(v4) > 0 {
                    v4[0] = 127
                }
                v1 = v1
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v23 := GenS5{F0: ("foo bar" + "hello")}
    v1 += (v1 + (v1 % ((v1) | 1)))
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 bool = ((!(int16(v1) >= 218)) && true)
            v3 = fi24
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS6{F0: (false && false), F1: 2147483647, F2: (^(v1))}
    fmt.Println(v3)
    v4 = append(v4, v3)
    if len(v4) > 0 {
        v4[0] = (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
