package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
func genFn0() int64 {
    var v5 uint = 18446744073709551615
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: "hello"}
    v8 := []int{-642}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    v3 := []uint32{679}
    var v4 int32 = (v2 | 42)
    v2 = 1
    _ = len(v3)
    v9 := v1[(("\n\t" + "\n\t") + "foo bar")]
    v10 := GenS1{F0: (("hello" + "") + "")}
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println(v2, v9)
    v11 := v1["hello"]
    v12 := GenS2{F0: 2}
    v13 := v1[(("hello" + ("" + "")) + "")]
    for fi14 := 0; fi14 < 18; fi14++ {
        switch 55 {
        case -150:
            v3 = append(v3, 1)
        case 127:
            var v15 uint = 233
            _ = v15
        case -532:
            switch ("\n\t" + (("\n\t" + "\n\t") + ("hello" + (("hello" + "\n\t") + "foo bar")))) {
            case "\n\t":
                var v16 uint8 = 2
                _ = v16
            case "a":
                var v17 float64 = 1.5
                _ = v17
            }
        }
        fmt.Println(v2, v2, v4)
        v18 := GenS1{F0: ""}
        _ = fi14
        _ = v18
    }
    if false {
        if (true || (int16(v9) < int16(v2))) {
            v19 := v1[((("hello" + "\n\t") + "\n\t") + ("" + ("\n\t" + "")))]
            _ = v19
        }
        v20 := GenS1{F0: (("\n\t" + "a") + "foo bar")}
        v2 = v2
        _ = v20
    }
    v11 += v9
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 2147483647)
    v9 -= v9
    v21 := v1["\n\t"]
    v9 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
