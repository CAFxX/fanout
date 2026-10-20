package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 uint }

func main() {
    var v1 uint16 = 2
    var v2 int = 0
    v3 := []int{v2, v2, (v2 + (v2 >> ((146) & 63))), v2}
    v4 := GenS1{F0: -128, F1: "\n\t"}
    v2 -= v2
    switch "a" {
    case "a":
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 38, F1: (("" + "a") + "")}
            switch (v2 | v2) {
            case 438:
                v3 = append(v3, fi5)
            }
            if len(v3) > 0 {
                v3[0] = (v2 ^ (-(v2)))
            }
            _ = fi5
            _ = v6
        }
    case "":
        v7 := GenS1{F0: -115, F1: (("foo bar" + "") + "a")}
        _ = v7
    }
    var v8 int32 = -323
    var v9 uint32 = 2147483648
    for fi10 := 0; fi10 < 12; fi10++ {
        v8 -= v8
        fmt.Println(v1, fi10, v9)
        if (1 <= -9223372036854775808) {
            if ((!("a" <= "foo bar")) && false) {
                v11 := GenS1{F0: -9, F1: "\n\t"}
                _ = v11
            } else {
                if ((v2 & fi10) != (-(fi10))) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 int64 = 35
                        v14 := GenS2{F0: 42, F1: 42}
                        if true {
                            if true {
                                v2 -= (fi12 + 1)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v1 -= 103
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v8 -= -1
                    if ("" <= "foo bar") {
                        v8 = -571
                        v1 -= (v1 - 0)
                    } else {
                        if true {
                            var v15 int = -9223372036854775808
                            _ = v15
                        }
                    }
                }
                var v16 uint8 = 2
                v1 -= 0
                _ = v16
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 := GenS1{F0: -1, F1: (("foo bar" + "a") + "a")}
            _ = v17
        }
        _ = fi10
    }
    if true {
        v9 -= 580
        for fi18 := 0; fi18 < 17; fi18++ {
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            switch fi18 {
            case 607:
                v1 -= 970
            }
            v19 := GenS3{F0: 42, F1: (v1 << ((0) & 63)), F2: 616}
            _ = fi18
            _ = v19
        }
    }
    if true {
        fmt.Println(v9, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
