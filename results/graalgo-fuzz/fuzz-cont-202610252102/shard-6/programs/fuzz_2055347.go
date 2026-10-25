package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1.5
    v3 := make(map[string]string)
    v4 := v1[("a" + ("" + "a"))]
    var v5 int = -1
    switch (("a" + (("hello" + ("hello" + "a")) + "foo bar")) + "foo bar") {
    case "\n\t":
        v1["hello"] = v4
    case "a":
        v4 = v4
    }
    v4 = v4
    if true {
        if true {
            v2 = 1e10
            var v6 uint8 = 42
            _ = v6
        } else {
            switch "foo bar" {
            case "":
                v5 = v5
            case "\n\t":
                var v7 float64 = v2
                _ = v7
            }
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 2, F1: 4294967295}
            var v10 uint16 = 836
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v4 = (fi11 > -9223372036854775808)
        _ = fi11
    }
    if v4 {
        v4 = (885 >= 1)
        v4 = false
        v12 := GenS2{F0: 745}
        _ = v12
    }
    var v13 uint64 = 165
    if v4 {
        v2 += ((v2 / float64(-2.25)) + -2.25)
        switch "\n\t" {
        case "":
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 uint32 = 2147483648
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        }
        if (0 < -2147483648) {
            v17 := GenS2{F0: 65535}
            _ = v17
        } else {
            v18 := GenS2{F0: 146}
            _ = v18
        }
    } else {
        if (!(!(2147483647 > 42))) {
            if v4 {
                v4 = true
                if false {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v13 -= v13
                        v20 := GenS3{F0: 0}
                        _ = fi19
                        _ = v20
                    }
                    var v21 string = ("" + ("\n\t" + "foo bar"))
                    _ = v21
                }
            }
            v22 := GenS4{F0: 0, F1: 4294967295}
            for fi23 := 0; fi23 < 11; fi23++ {
                v5 += v5
                v13 -= 559
                _ = fi23
            }
            _ = v22
        } else {
            v24 := GenS2{F0: 255}
            v5 = 310
            _ = v24
        }
    }
    v25 := GenS3{F0: -991}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
