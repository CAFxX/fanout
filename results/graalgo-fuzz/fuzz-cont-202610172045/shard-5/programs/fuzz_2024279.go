package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 int16, p1 uint16, p2 uint) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 = 0
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int16 = (p0 | (^(p0)))
            p2 = p2
            var v10 bool = (2147483648 > 42)
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{127, (fi7 + (fi7 << ((p2) & 63))), fi7, (fi7 / (((fi7 ^ (^(fi7)))) | 1))}
        _ = fi7
        _ = v11
    }
    v12 := GenS1{F0: 4294967295, F1: 111}
    fmt.Println(p1, p2, p1)
    p0 = p0
    v13 := GenS1{F0: -9223372036854775808, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}
func genFn1() string {
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 float64 = 1e10
            _ = fi15
            _ = v16
        }
        fmt.Println(fi14)
        fmt.Println(fi14)
        _ = fi14
    }
    var v17 uint64 = 1
    switch ((("" + "a") + "\n\t") + "a") {
    case "\n\t":
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(fi18)
            if (!false) {
                if false {
                    v19 := GenS2{F0: 0}
                    v17 -= 1
                    _ = v19
                }
                fmt.Println(fi18, v17)
                v20 := GenS2{F0: 9223372036854775807}
                _ = v20
            }
            _ = fi18
        }
    }
    v21 := []int{-291, 2147483647, 128}
    _ = v17
    _ = v21
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 string = (("foo bar" + (("" + ("hello" + "a")) + (("\n\t" + "foo bar") + (("\n\t" + "a") + "foo bar")))) + ("\n\t" + "foo bar"))
    v4 := make(map[int]string)
    var v5 uint32 = 2147483647
    var v6 bool = (4294967295 == 796)
    fmt.Println(v3)
    var v22 uint8 = 71
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int = (fi23 << ((0) & 63))
        _ = fi23
        _ = v24
    }
    var v25 uint8 = (v22 & (v22 >> ((937) & 63)))
    v25 += (v25 << ((uint(v5)) & 63))
    fmt.Println(v22)
    v6 = v6
    switch 32767 {
    case 2147483647:
        v26 := v1[v3]
        _ = v26
    case -1:
        if (v6 && true) {
            for fi27 := 0; fi27 < 5; fi27++ {
                if v6 {
                    for fi28 := 0; fi28 < 10; fi28++ {
                        fmt.Println(v6, fi28, v25)
                        v3 = v3
                        v29 := v2[v5]
                        _ = fi28
                        _ = v29
                    }
                    v30 := GenS3{F0: v5, F1: -1, F2: 163}
                    v31 := GenS2{F0: 9223372036854775807}
                    _ = v30
                    _ = v31
                }
                _ = fi27
            }
            v6 = (v3 > "hello")
        } else {
            v32 := []int{128, 128, 0, 32768}
            v3 = v3
            if (32767 == 818) {
                for fi33 := 0; fi33 < 13; fi33++ {
                    var v34 uint8 = (v22 | v25)
                    var v35 uint32 = v5
                    switch (^(fi33)) {
                    case 0:
                        fmt.Println(v35, v6, v6)
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 bool = (true && true)
                var v37 uint8 = (^(v25))
                _ = v36
                _ = v37
            }
            _ = v32
        }
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(fi38, fi38, v6)
        _ = fi38
    }
    var v39 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
