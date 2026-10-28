package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 int8, p1 uint16) int8 {
    v6 := []int{128, 128, 32768}
    v7 := GenS1{F0: 2147483647, F1: 127}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 string = ((("" + ("" + "foo bar")) + "\n\t") + ("\n\t" + "foo bar"))
    var v2 uint = 4294967295
    var v3 uint = v2
    v4 := []int{}
    v5 := []uint32{248}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: (v1 + "foo bar"), F1: 32767, F2: 0}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 float32 = 3.14159
    for fi12 := 0; fi12 < 16; fi12++ {
        v2 -= 50
        var v13 int = fi12
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint8 = 71
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if (!(847 > 42)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (v1 > "foo bar") {
                if len(v4) > 0 {
                    v4[0] = 127
                }
            }
            _ = fi16
        }
    }
    v3 = 18446744073709551615
    _ = len(v4)
    switch (v1 + ("\n\t" + ("\n\t" + "a"))) {
    case "a":
        v4 = append(v4, 1)
    }
    v11 = 1.5
    if (true && (!true)) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = (("a" + v1) + ("hello" + ((((v1 + "a") + v1) + (v1 + v1)) + v1)))
    } else {
        switch 42 {
        case 33:
            var v17 int8 = -128
            _ = v17
        case 42:
            fmt.Println(v1, v11)
        case 280:
            v18 := GenS1{F0: 526, F1: 2}
            _ = v18
        }
    }
    v2 = v2
    v19 := GenS1{F0: -148, F1: 2}
    v2 = v2
    var v20 int32 = 42
    v2 += 65535
    if false {
        v21 := GenS2{F0: "hello", F1: 1, F2: 42}
        _ = v21
    } else {
        v3 = v2
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        _ = len(v5)
        var v23 bool = true
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
