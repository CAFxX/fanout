package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"foo bar", ("hello" + ("" + ((("" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) + "a") + "a")))}
    var v3 uint8 = 0
    var v4 bool = (858 == 0)
    var v5 int32 = (v1 | (v1 - -1))
    var v6 int = -658
    v7 := GenS1{F0: 2}
    v1 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v4, v5, v4)
        _ = len(v2)
        _ = fi8
    }
    fmt.Println(v5)
    var v9 uint16 = 2
    fmt.Println(v9, v5, v1)
    for fi10 := 0; fi10 < 12; fi10++ {
        _ = len(v2)
        _ = fi10
    }
    switch 656 {
    case 127:
        var v11 uint16 = 152
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (fi12 / (((-(fi12))) | 1)) {
            case -894:
                var v13 uint = 42
                _ = v13
            case 32768:
                var v14 string = "foo bar"
                _ = v14
            case 1:
                v9 -= v9
            default:
                if (false || (v4 || (false && (("\n\t" + "foo bar") <= "")))) {
                    fmt.Println(v9, v6)
                    var v15 int16 = -32768
                    _ = v15
                } else {
                    var v16 int = int(v3)
                    fmt.Println(v3)
                    v2 = append(v2, "")
                    _ = v16
                }
            }
            v17 := GenS1{F0: 270}
            _ = fi12
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            v3 -= 255
            v9 = (v9 << ((255) & 63))
            _ = fi18
        }
    } else {
        var v19 int = (v6 + int(v3))
        v1 = (v1 + (v1 & -1))
        _ = v19
    }
    switch 0 {
    case 2147483647:
        v20 := GenS1{F0: 4294967295}
        _ = v20
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v21 uint16 = (^(v9))
    var v22 float32 = -2.25
    v9 = (v21 | (v21 | v9))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
