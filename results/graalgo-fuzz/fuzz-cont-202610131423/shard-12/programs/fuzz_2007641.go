package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64) string {
    p0 = p0
    switch 533 {
    case 9223372036854775807:
        v4 := []int{-2147483648, 4294967295, 42}
        _ = v4
    case 4294967295:
        var v5 uint16 = 0
        _ = v5
    case 139:
        p0 = p0
    default:
        switch -489 {
        case -9223372036854775808:
            var v6 bool = (-2.25 >= -2.25)
            _ = v6
        }
    }
    var v7 bool = false
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 42
    v3 := []string{("\n\t" + (("foo bar" + "foo bar") + ("hello" + ((("a" + "a") + "") + "")))), "foo bar", "hello"}
    if false {
        v3 = append(v3, (("" + "hello") + ("\n\t" + "foo bar")))
        switch 32768 {
        case -2147483648:
            v8 := GenS1{F0: 903}
            _ = v8
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v9 bool = false
    var v10 string = ""
    _ = len(v3)
    v10 = (((("\n\t" + ((v10 + v10) + v10)) + "foo bar") + (v10 + "\n\t")) + ((v10 + v10) + v10))
    v3 = append(v3, ("\n\t" + "\n\t"))
    for fi11 := 0; fi11 < 8; fi11++ {
        switch (-(fi11)) {
        case 32767:
            v2 = v2
        case -2147483648:
            if v9 {
                var v12 uint32 = (v2 ^ v2)
                _ = v12
            } else {
                v13 := GenS2{F0: 32767}
                v10 = v10
                _ = v13
            }
        }
        _ = fi11
    }
    v14 := GenS1{F0: 42}
    var v15 int = 128
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: 42}
            fmt.Println(v10, v1, v2)
            _ = fi17
            _ = v18
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v10 = v10
        _ = fi16
    }
    var v19 uint16 = 1
    var v20 int64 = -580
    var v21 int64 = (v20 >> ((66) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
