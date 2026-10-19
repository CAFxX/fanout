package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float32 = ((1e10 * (1.5 - float32(1.5))) / 0.1)
    var v2 uint = 65535
    var v3 int32 = -599
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            fmt.Println(v3, v3, v2)
            if (!(-115 >= -48)) {
                for fi5 := 0; fi5 < 18; fi5++ {
                    fmt.Println(fi4)
                    _ = fi5
                }
            }
            var v6 string = "hello"
            _ = v6
        }
        switch (fi4 & (^(fi4))) {
        case 0:
            switch (fi4 + 4294967295) {
            case 2147483647:
                fmt.Println(v1)
            default:
                v2 = v2
            }
        }
        _ = fi4
    }
    v1 += v1
    if true {
        v7 := GenS1{F0: -461}
        v3 -= (v3 & (v3 + (v3 / (((v3 + (v3 + (v3 - (v3 ^ v3))))) | 1))))
        var v8 uint32 = 10
        _ = v7
        _ = v8
    } else {
        v9 := []int{1, 37}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v9
    }
    fmt.Println(v2)
    v10 := GenS1{F0: 290}
    v11 := []int{1}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 uint16 = 65535
    var v13 int32 = (^(v3))
    for fi14 := 0; fi14 < 18; fi14++ {
        if true {
            v11 = append(v11, -578)
            if ("hello" > ("foo bar" + "hello")) {
                if true {
                    var v15 int64 = 0
                    _ = v15
                }
            }
            v12 -= v12
        } else {
            var v16 int = (fi14 ^ fi14)
            if len(v11) > 0 {
                v11[0] = (fi14 * (-(fi14)))
            }
            _ = v16
        }
        var v17 int8 = -13
        _ = fi14
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
