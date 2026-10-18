package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 uint64, p2 int8) string {
    fmt.Println(p0, p1)
    var v7 uint32 = 65535
    v7 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint, p1 uint16) uint8 {
    p0 = p0
    if (42 > 556) {
        fmt.Println(p0)
        v8 := []int{-9223372036854775808, 42, -96, -817}
        _ = v8
    } else {
        v9 := GenS1{F0: 654}
        if false {
            for fi10 := 0; fi10 < 13; fi10++ {
                switch (fi10 >> ((p0) & 63)) {
                case 4294967295:
                    var v11 uint32 = 742
                    _ = v11
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    p0 = p0
                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                    case "foo bar":
                        v13 := GenS1{F0: 175}
                        _ = v13
                    case "hello":
                        fmt.Println(p0)
                    case "\n\t":
                        p0 -= p0
                    }
                    v14 := []int{32767}
                    _ = fi12
                    _ = v14
                }
                var v15 int16 = 627
                _ = fi10
                _ = v15
            }
            var v16 int = -2147483648
            _ = v16
        } else {
            fmt.Println(p0, p1)
            var v17 uint = (p0 >> ((p0) & 63))
            _ = v17
        }
        _ = v9
    }
    var v18 uint64 = 475
    if (0 == 204) {
        v19 := []int{-918}
        fmt.Println(p0, p1, p1)
        _ = v19
    }
    if true {
        switch -294 {
        case 1:
            if true {
                switch -66 {
                case -2147483648:
                    var v20 int8 = -128
                    _ = v20
                case 264:
                    if (!(0 < 2147483647)) {
                        v21 := []int{42, 127}
                        for fi22 := 0; fi22 < 15; fi22++ {
                            if len(v21) > 0 {
                                v21[0] = 644
                            }
                            _ = fi22
                        }
                        _ = v21
                    }
                default:
                    v23 := GenS2{F0: -101}
                    _ = v23
                }
            } else {
                v24 := []int{2147483647}
                fmt.Println(v18, p0)
                _ = v24
            }
        }
        v18 = 4294967295
        if (603 <= 2147483648) {
            for fi25 := 0; fi25 < 17; fi25++ {
                p0 = (p0 * 286)
                var v26 string = ((((("\n\t" + (("" + "hello") + "hello")) + ("\n\t" + "hello")) + "a") + ("\n\t" + "")) + (("" + "\n\t") + ("" + "a")))
                _ = fi25
                _ = v26
            }
            v27 := GenS2{F0: 0}
            _ = v27
        }
    } else {
        v18 += (^(v18))
        for fi28 := 0; fi28 < 1; fi28++ {
            p0 -= (p0 * 42)
            _ = fi28
        }
        switch (("\n\t" + ("hello" + "\n\t")) + (("foo bar" + "a") + ("hello" + ""))) {
        case "a":
            var v29 uint16 = 65535
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = v18
    return 255
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 uint16 = 0
    var v4 uint16 = (v3 << ((2) & 63))
    var v5 uint8 = 240
    var v6 uint = 1
    fmt.Println(v4)
    v5 = (v5 * v5)
    if (!true) {
        v30 := []int{int(v6), 4294967295, 1}
        for fi31 := 0; fi31 < 12; fi31++ {
            v5 = (v5 % ((v5) | 1))
            _ = fi31
        }
        _ = v30
    } else {
        switch ("" + "a") {
        case "":
            var v32 uint = v6
            _ = v32
        default:
            v4 = v4
        }
        var v33 bool = true
        _ = v33
    }
    v34 := []int{-35, 241, 1, 0}
    v5 -= v5
    switch 509 {
    case 989:
        fmt.Println(v6, v1)
    case 1:
        var v35 int16 = 1000
        _ = v35
    }
    v4 -= v3
    var v36 uint32 = 65535
    for fi37 := 0; fi37 < 17; fi37++ {
        v5 -= v5
        v6 = v6
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 string = "a"
            v36 -= (v36 ^ 82)
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
