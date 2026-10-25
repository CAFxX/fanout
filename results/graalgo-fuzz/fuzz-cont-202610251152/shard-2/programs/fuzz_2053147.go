package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{"a", "", "", ("foo bar" + "foo bar")}
    var v2 string = ("foo bar" + ("" + "a"))
    var v3 uint8 = 82
    v4 := []uint32{382, 4294967295, 42}
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v2 = ((v2 + "a") + v2)
            _ = fi7
        }
        _ = fi6
    }
    fmt.Println(v3)
    v3 += (v3 % (((v3 >> ((150) & 63))) | 1))
    v1 = append(v1, ("foo bar" + v2))
    if false {
        var v8 int64 = -1
        if (!(4294967295 < 15)) {
            if false {
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: fi9}
                    _ = fi9
                    _ = v10
                }
            } else {
                v8 = (v8 + v8)
                v8 = (v8 / ((v8) | 1))
                fmt.Println(v8, v2)
            }
            switch (v2 + "hello") {
            case "hello":
                var v11 int16 = 1
                _ = v11
            }
        }
        _ = v8
    } else {
        var v12 string = ("foo bar" + (((((v2 + ((((("hello" + v2) + ((("\n\t" + (v2 + "foo bar")) + "\n\t") + ("\n\t" + "\n\t"))) + ("a" + ("\n\t" + v2))) + v2) + ("\n\t" + v2))) + (v2 + v2)) + "hello") + v2) + v2))
        _ = v12
    }
    v2 = v2
    fmt.Println(v3, v2)
    v5[-1] = (true || true)
    _ = len(v5)
    var v13 string = ""
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, ("a" + (v13 + v2)))
    v14 := GenS1{F0: -781}
    v15 := GenS1{F0: -131}
    var v16 uint16 = 1
    if true {
        v3 -= (v3 << ((2) & 63))
        fmt.Println(v13)
        var v17 uint = uint(v16)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 3; fi18++ {
            v3 += (v3 + (v3 >> ((2) & 63)))
            _ = fi18
        }
        switch -9223372036854775808 {
        case 34:
            _ = len(v4)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
