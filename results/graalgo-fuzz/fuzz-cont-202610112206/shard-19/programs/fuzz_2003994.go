package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 int16 = -478
    v3 += v3
    var v4 float32 = ((3.14159 - 3.14159) * -2.25)
    fmt.Println(v4, v4)
    fmt.Println(v4, v4)
    v5 := GenS1{F0: -1}
    if false {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 13; fi7++ {
                v1 = append(v1, "foo bar")
                fmt.Println(fi7)
                v4 -= (v4 / v4)
                _ = fi7
            }
            if len(v1) > 0 {
                v1[0] = ("a" + "a")
            }
            _ = fi6
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                v4 = v4
            }
            _ = fi8
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v9 string = "hello"
        var v10 uint64 = 4294967295
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v3)
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS2{F0: 1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            switch (-(fi14)) {
            case 2147483647:
                v15 := GenS1{F0: 0}
                _ = v15
            }
            if false {
                v3 = v3
                v16 := GenS1{F0: 1}
                _ = v16
            }
            _ = fi14
        }
        _ = fi11
    }
    var v17 float64 = (((float64(3.14159) * (0.0 / ((0.0 + 3.14159) + 0.1))) - -2.25) * -2.25)
    var v18 float32 = v4
    if (42 > 0) {
        for fi19 := 0; fi19 < 10; fi19++ {
            _ = len(v1)
            var v20 bool = (2 < 640)
            _ = fi19
            _ = v20
        }
        var v21 float32 = ((v18 * v4) * (0.1 * ((3.14159 - v18) - v4)))
        v3 += v3
        _ = v21
    } else {
        var v22 float64 = v17
        _ = v22
    }
    switch 4294967295 {
    case 127:
        fmt.Println(v18, v18, v3)
    case 194:
        var v23 bool = (true && (false && ((v3 - 994) < v3)))
        _ = v23
    }
    v24 := GenS1{F0: 0}
    switch 2147483647 {
    case 1:
        if (91 < 1) {
            var v25 float64 = v17
            _ = v25
        }
    case 17:
        v17 -= -2.25
    case 0:
        v4 += (1.5 * 0.1)
    }
    var v26 bool = ((0 == 0) || false)
    var v27 uint16 = 0
    fmt.Println(v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
