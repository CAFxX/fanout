package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{2147483648, 42}
    var v3 int = -2147483648
    var v4 int64 = -599
    var v5 uint8 = 0
    var v6 string = ("hello" + ("a" + "a"))
    var v7 bool = (797 != 42)
    var v8 int64 = -186
    v6 = "hello"
    var v9 uint16 = 32767
    var v10 uint64 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 := GenS1{F0: (v1 >> ((v1) & 63))}
    if v7 {
        for fi12 := 0; fi12 < 9; fi12++ {
            if v7 {
                var v13 string = v6
                fmt.Println(v4)
                var v14 float64 = -2.25
                _ = v13
                _ = v14
            } else {
                v10 += v10
                var v15 int = (^(v3))
                _ = v15
            }
            if len(v2) > 0 {
                v2[0] = 0
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS2{F0: (v5 << ((470) & 63))}
                v9 -= v9
                var v18 int64 = v4
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi12
        }
        var v19 int = v3
        _ = v19
    } else {
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 string = ""
            v22 := GenS1{F0: 18446744073709551615}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v2 = append(v2, 2)
    fmt.Println(v4)
    switch (("" + ((("foo bar" + ("\n\t" + "foo bar")) + ((v6 + v6) + v6)) + (v6 + "foo bar"))) + "") {
    case "a":
        switch "\n\t" {
        case "hello":
            var v23 int32 = -2147483648
            _ = v23
        case "a":
            fmt.Println(v8, v7, v10)
        case "foo bar":
            if (!true) {
                var v24 bool = (-1 >= -858)
                var v25 uint64 = (v10 % ((v10) | 1))
                var v26 float64 = -2.25
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 int8 = 127
                _ = v27
            }
        }
    case "":
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 int8 = 120
            var v30 int32 = int32(v8)
            _ = fi28
            _ = v29
            _ = v30
        }
    case "\n\t":
        switch (v6 + ("a" + (v6 + v6))) {
        case "":
            if v7 {
                if false {
                    v5 = (v5 & 139)
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                }
                v31 := GenS2{F0: v5}
                _ = v31
            } else {
                var v32 uint = v1
                _ = v32
            }
        }
    default:
        for fi33 := 0; fi33 < 4; fi33++ {
            for fi34 := 0; fi34 < 2; fi34++ {
                if v7 {
                    fmt.Println(v9)
                }
                var v35 bool = v7
                var v36 float32 = 1e10
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
